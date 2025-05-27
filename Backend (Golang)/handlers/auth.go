package handlers

import (
	"Go-Kurs/config"
	"Go-Kurs/logger"
	"Go-Kurs/models"
	"Go-Kurs/repository"
	"Go-Kurs/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"

	log "github.com/sirupsen/logrus"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

// Register создаёт нового пользователя.
// @Summary Регистрация пользователя
// @Description Регистрация нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 409 {object} map[string]interface{} "Логин уже существует"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Name       string `json:"name" validate:"required,min=2,max=20"`
		Surname    string `json:"surname" validate:"required,min=2,max=20"`
		Patronymic string `json:"patronymic" validate:"max=20"`
		Telephone  string `json:"telephone" validate:"required"`
		Login      string `json:"login" validate:"required,min=3,max=100"`
		Password   string `json:"password" validate:"required,min=6,max=100"`
	}

	logger.Log.Info("Register endpoint called")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to read request body")
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	logger.Log.WithField("request_body", string(body)).Debug("Raw request body")

	if err := json.Unmarshal(body, &user); err != nil {
		logger.Log.WithError(err).Error("Failed to decode registration request")
		http.Error(w, fmt.Sprintf("Invalid request format: %v", err), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logger.Log.WithError(err).Error("Validation failed")

		var errorMsg string
		for _, err := range err.(validator.ValidationErrors) {
			errorMsg += fmt.Sprintf("Field '%s' failed validation (%s); ", err.Field(), err.Tag())
		}

		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	dbUser := models.User{
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Telephone:  user.Telephone,
		Login:      user.Login,
		Password:   user.Password,
	}

	exists, err := h.userRepo.IsLoginExists(dbUser.Login)
	if err != nil {
		logger.Log.WithError(err).Error("Database error during login check")
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		logger.Log.WithField("login", dbUser.Login).Warn("Login already exists")
		http.Error(w, "Login already exists", http.StatusConflict)
		return
	}

	err = h.userRepo.CreateUser(&dbUser)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to create user")
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	token, err := dbUser.GenerateJWT(config.AppConfig.App.JWTSecret)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to generate token")
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	logger.Log.WithField("user_id", dbUser.ID).Info("User registered successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "User registered successfully",
		"token":   token,
		"user": map[string]interface{}{
			"id":      dbUser.ID,
			"name":    dbUser.Name,
			"surname": dbUser.Surname,
			"login":   dbUser.Login,
		},
	})
}

// Login авторизует пользователя.
// @Summary Авторизация пользователя
// @Description Авторизация пользователя по логину и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Логин и пароль"
// @Success 200 {object} map[string]interface{} "Успешный ответ с токеном"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 401 {object} map[string]interface{} "Неверный логин или пароль"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Login endpoint called")
	var credentials models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		logger.Log.WithError(err).Error("Failed to decode login request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByLogin(credentials.Login)
	if err != nil {
		logger.Log.WithError(err).Error("Database error during authentication")
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if user == nil || !user.CheckPassword(credentials.Password) {
		logger.Log.WithField("login", credentials.Login).Warn("Invalid login attempt")
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	token, err := user.GenerateJWT(config.AppConfig.App.JWTSecret)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to generate token")
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	logger.Log.WithField("user_id", user.ID).Info("User logged in successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Login successful",
		"token":   token,
		"user": map[string]interface{}{
			"id":      user.ID,
			"name":    user.Name,
			"surname": user.Surname,
			"login":   credentials.Login,
		},
	})
}

// AddFavorite добавляет товар в избранное.
// @Summary Добавить товар в избранное
// @Description Добавление товара в список избранного для пользователя
// @Tags favorites
// @Accept json
// @Produce json
// @Param request body models.Request true "ID пользователя и ID товара"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 404 {object} map[string]interface{} "Товар или пользователь не найдены"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/favorites/add [post]
func (h *AuthHandler) AddFavorite(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("AddFavorite endpoint called")
	w.Header().Set("Content-Type", "application/json")

	var request models.Request // Используем новый тип
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Log.WithError(err).Error("Invalid request format in AddFavorite")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid request format: " + err.Error(),
		})
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Debug("Adding product to favorites")

	if err := h.userRepo.AddToFavorites(request.UserID, request.ProductID); err != nil {
		status := http.StatusInternalServerError
		errorMsg := err.Error()

		if strings.Contains(errorMsg, "does not exist") {
			status = http.StatusNotFound
			logger.Log.WithError(err).Warn("Product or user not found")
		} else {
			logger.Log.WithError(err).Error("Failed to add to favorites")
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   errorMsg,
		})
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Info("Product added to favorites successfully")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product successfully added to favorites",
	})
}

// RemoveFavorite удаляет товар из избранного.
// @Summary Удалить товар из избранного
// @Description Удаление товара из списка избранного для пользователя
// @Tags favorites
// @Accept json
// @Produce json
// @Param request body models.Request true "ID пользователя и ID товара"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 404 {object} map[string]interface{} "Товар или пользователь не найдены"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/favorites/remove [post]
func (h *AuthHandler) RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	log.Println("RemoveFavorite endpoint hit")
	logger.Log.Info("RemoveFavorite endpoint called")

	var request models.Request

	body, _ := io.ReadAll(r.Body)
	log.Printf("Request body: %s", string(body))
	logger.Log.WithField("request_body", string(body)).Debug("RemoveFavorite request body")
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Decode error: %v", err)
		logger.Log.WithError(err).Error("Failed to decode RemoveFavorite request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Removing favorite: user_id=%d, product_id=%d", request.UserID, request.ProductID)
	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Info("Removing product from favorites")

	if err := h.userRepo.RemoveFromFavorites(request.UserID, request.ProductID); err != nil {
		log.Printf("Remove error: %v", err)
		logger.Log.WithError(err).Error("Failed to remove from favorites")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Info("Product removed from favorites")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Removed from favorites",
	})
}

// GetFavorites возвращает список избранных товаров пользователя.
// @Summary Получить избранное пользователя
// @Description Получение списка избранных товаров для данного пользователя
// @Tags favorites
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Success 200 {array} models.Product "Список избранных товаров"
// @Failure 400 {object} map[string]interface{} "Неверный ID пользователя"
// @Failure 404 {object} map[string]interface{} "Избранное не найдено"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/favorites [get]
func (h *AuthHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("GetFavorites endpoint called")
	//userID := r.Context().Value("userID").(int)
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Log.WithError(err).Error("Invalid user ID in GetFavorites")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("user_id", userID).Debug("Getting favorites for user")

	products, err := h.userRepo.GetFavorites(userID)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to get favorites")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id": userID,
		"count":   len(products),
	}).Info("Favorites retrieved successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// AddToBasket добавляет товар в корзину.
// @Summary Добавить товар в корзину
// @Description Добавление товара в корзину для пользователя
// @Tags basket
// @Accept json
// @Produce json
// @Param request body models.Request true "ID пользователя и ID товара"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 404 {object} map[string]interface{} "Товар или пользователь не найдены"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/basket/add [post]
func (h *AuthHandler) AddToBasket(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("AddToBasket endpoint called")

	var request models.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Log.WithError(err).Error("Invalid request format in AddToBasket")

		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Debug("Adding product to basket")

	if err := h.userRepo.AddToBasket(request.UserID, request.ProductID); err != nil {
		logger.Log.WithError(err).Error("Failed to add to basket")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Info("Product added to basket successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product added to basket",
	})
}

// RemoveFromBasket удаляет товар из корзины.
// @Summary Удалить товар из корзины
// @Description Удаление товара из корзины для пользователя
// @Tags basket
// @Accept json
// @Produce json
// @Param request body models.Request true "ID пользователя и ID товара"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 404 {object} map[string]interface{} "Товар или пользователь не найдены"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/basket/remove [post]
func (h *AuthHandler) RemoveFromBasket(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("RemoveFromBasket endpoint called")

	var request models.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Log.WithError(err).Error("Invalid request format in RemoveFromBasket")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Debug("Removing product from basket")

	if err := h.userRepo.RemoveFromBasket(request.UserID, request.ProductID); err != nil {
		logger.Log.WithError(err).Error("Failed to remove from basket")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":    request.UserID,
		"product_id": request.ProductID,
	}).Info("Product removed from basket successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product removed from basket",
	})
}

// GetBasket возвращает корзину пользователя.
// @Summary Получить корзину пользователя
// @Description Получение товаров в корзине для данного пользователя
// @Tags basket
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Success 200 {array} models.Product "Список товаров в корзине"
// @Failure 400 {object} map[string]interface{} "Неверный ID пользователя"
// @Failure 404 {object} map[string]interface{} "Корзина не найдена"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/basket [get]
func (h *AuthHandler) GetBasket(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("GetBasket endpoint called")

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Log.WithError(err).Error("Invalid user ID in GetBasket")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("user_id", userID).Debug("Getting basket for user")
	products, err := h.userRepo.GetBasket(userID)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to get basket")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id": userID,
		"count":   len(products),
	}).Info("Basket retrieved successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CreateOrder создаёт новый заказ.
// @Summary Создание нового заказа
// @Description Создание заказа на основе выбранных товаров и данных покупателя
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.OrderRequest true "Данные заказа"
// @Success 201 {object} map[string]interface{} "Успешный ответ с ID заказа"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/orders [post]
func (h *AuthHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("CreateOrder endpoint called")
	w.Header().Set("Content-Type", "application/json")

	log.Println("Incoming order request")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to read request body")
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to read request"})
		return
	}
	defer r.Body.Close()

	logger.Log.WithField("request_body", string(body)).Debug("CreateOrder request")
	log.Printf("Request body: %s", string(body))

	var request models.OrderRequest

	if err := json.Unmarshal(body, &request); err != nil {
		logger.Log.WithError(err).Error("Failed to decode JSON in CreateOrder")
		log.Printf("JSON decode error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	logger.Log.WithField("request", request).Debug("Parsed CreateOrder request")
	log.Printf("Parsed request: %+v", request)

	if request.CustomerID == 0 || len(request.Products) == 0 {
		logger.Log.Warn("Missing required fields in CreateOrder")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing required fields"})
		return
	}

	tx, err := h.userRepo.GetDB().Begin()
	if err != nil {
		logger.Log.WithError(err).Error("Failed to begin transaction")
		log.Printf("Transaction begin error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Log.WithError(err).Error("Transaction rolled back")
		}
	}()

	var addressID int
	err = tx.QueryRow(`
        INSERT INTO addresses (customer_id, country, city, street, house, apartment)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `, request.CustomerID, request.Address.Country, request.Address.City,
		request.Address.Street, request.Address.House, request.Address.Apartment).Scan(&addressID)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to create address")
		log.Printf("Address insert error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create address"})
		return
	}

	var orderID int
	err = tx.QueryRow(`
        INSERT INTO orders (customer_id, order_price, card_number, order_date)
        VALUES ($1, $2, $3, NOW())
        RETURNING id
    `, request.CustomerID, request.OrderPrice, request.CardNumber).Scan(&orderID)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to create order")
		log.Printf("Order insert error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create order"})
		return
	}

	for _, productID := range request.Products {
		_, err = tx.Exec(`
            INSERT INTO product_order (product_id, order_id)
            VALUES ($1, $2)
        `, productID, orderID)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to add product to order")
			log.Printf("Product_order insert error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to add products to order"})
			return
		}
	}

	_, err = tx.Exec(`
        INSERT INTO deliveries (order_id, address_id, status, expected_receive_date)
        VALUES ($1, $2, 'Заказ собирается', NOW() + interval '4 days')
    `, orderID, addressID)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to create delivery")
		log.Printf("Delivery insert error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create delivery"})
		return
	}

	_, err = tx.Exec(`
        DELETE FROM product_baskets pb
        USING baskets b
        WHERE pb.baskets_id = b.id AND b.customer_id = $1
    `, request.CustomerID)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to clear basket")
		log.Printf("Basket clear error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to clear basket"})
		return
	}

	if err = tx.Commit(); err != nil {
		logger.Log.WithError(err).Error("Failed to commit transaction")
		log.Printf("Transaction commit error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	logger.Log.WithFields(log.Fields{
		"order_id":    orderID,
		"customer_id": request.CustomerID,
		"item_count":  len(request.Products),
	}).Info("Order created successfully")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Order created successfully",
		"order_id": orderID,
	})
}

// ClearBasket очищает корзину пользователя.
// @Summary Очистить корзину пользователя
// @Description Удаление всех товаров из корзины для данного пользователя
// @Tags basket
// @Accept json
// @Produce json
// @Param request body models.Request2 true "ID пользователя"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]interface{} "Ошибка при валидации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/basket/clear [post]
func (h *AuthHandler) ClearBasket(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("ClearBasket endpoint called")

	var request models.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Log.WithError(err).Error("Invalid request format in ClearBasket")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("user_id", request.UserID).Debug("Clearing basket")

	result, err := h.userRepo.GetDB().Exec(`
        DELETE FROM product_baskets pb
        USING baskets b
        WHERE pb.baskets_id = b.id AND b.customer_id = $1
    `, request.UserID)

	if err != nil {
		logger.Log.WithError(err).Error("Failed to clear basket")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Log.WithError(err).Error("Failed to get rows affected")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithFields(log.Fields{
		"user_id":      request.UserID,
		"rows_deleted": rowsAffected,
	}).Info("Basket cleared successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Cleared %d items from basket", rowsAffected),
	})
}

func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.Log.Warn("Authorization header missing")
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseToken(authHeader, config.AppConfig.App.JWTSecret)
		if err != nil {
			logger.Log.WithError(err).Warn("Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
