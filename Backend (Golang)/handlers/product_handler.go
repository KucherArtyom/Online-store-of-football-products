package handlers

import (
	"Go-Kurs/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// GetAllProducts возвращает список всех товаров.
// @Summary Получение всех товаров
// @Description Получение списка всех товаров в магазине
// @Tags products
// @Produce json
// @Success 200 {array} models.Product "Список товаров"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByID возвращает товар по его ID.
// @Summary Получение товара по ID
// @Description Получение товара по его уникальному идентификатору
// @Tags products
// @Produce json
// @Param id path int true "ID товара"
// @Success 200 {object} models.Product "Товар"
// @Failure 400 {object} map[string]interface{} "Неверный ID"
// @Failure 404 {object} map[string]interface{} "Товар не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.repo.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

// GetProductsByCategory возвращает товары по категории.
// @Summary Получение товаров по категории
// @Description Получение товаров по идентификатору категории
// @Tags products
// @Produce json
// @Param category_id path int true "ID категории"
// @Success 200 {array} models.Product "Список товаров"
// @Failure 400 {object} map[string]interface{} "Неверный ID категории"
// @Failure 404 {object} map[string]interface{} "Товары не найдены"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /api/products/category/{category_id} [get]
func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	products, err := h.repo.GetProductsByCategory(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
