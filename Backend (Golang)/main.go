/*
package main

import (

	"Go-Kurs/handlers"
	"Go-Kurs/repository"
	"log"
	"net/http"





	"github.com/gorilla/mux"

)

	func enableCORS(router *mux.Router) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Разрешаем запросы с Vue-сервера разработки
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// Если это предварительный OPTIONS-запрос - завершаем обработку
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Передаем запрос дальше в роутер
			router.ServeHTTP(w, r)
		})
	}

	func main() {
		// Инициализация БД
		db, err := repository.NewPostgresDB()
		if err != nil {
			log.Fatalf("Failed to initialize DB: %v", err)
		}

		// Инициализация репозитория
		userRepo := repository.NewUserRepository(db)
		productRepo := repository.NewProductRepository(db)

		// Инициализация обработчиков
		authHandler := handlers.NewAuthHandler(userRepo)
		productHandler := handlers.NewProductHandler(productRepo)

		// Настройка маршрутов
		r := mux.NewRouter()
		r.HandleFunc("/api/register", authHandler.Register).Methods("POST", "OPTIONS")
		r.HandleFunc("/api/login", authHandler.Login).Methods("POST", "OPTIONS")
		// Настройка статики для production (если нужно)
		// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist")))

		r.HandleFunc("/api/products", productHandler.GetAllProducts).Methods("GET")
		r.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
		r.HandleFunc("/api/products/category/{category_id}", productHandler.GetProductsByCategory).Methods("GET")

		r.HandleFunc("/api/favorites/add", authHandler.AddFavorite).Methods("POST", "OPTIONS")
		r.HandleFunc("/api/favorites/remove", authHandler.RemoveFavorite).Methods("POST", "OPTIONS")
		r.HandleFunc("/api/favorites", authHandler.GetFavorites).Methods("GET")

		r.HandleFunc("/api/basket/add", authHandler.AddToBasket).Methods("POST", "OPTIONS")
		r.HandleFunc("/api/basket/remove", authHandler.RemoveFromBasket).Methods("POST", "OPTIONS")
		r.HandleFunc("/api/basket", authHandler.GetBasket).Methods("GET")

		r.HandleFunc("/api/orders", authHandler.CreateOrder).Methods("POST", "OPTIONS")
		r.HandleFunc("/api/basket/clear", authHandler.ClearBasket).Methods("POST", "OPTIONS")
		// Запуск сервера с CORS middleware
		log.Println("Server is running on port 8080...")
		log.Fatal(http.ListenAndServe(":8080", enableCORS(r)))
	}
*/

package main

import (
	"Go-Kurs/config"
	"Go-Kurs/handlers"
	"Go-Kurs/logger"
	"Go-Kurs/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func enableCORS(router *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		router.ServeHTTP(w, r)
	})
}

func main() {
	// Инициализация конфигурации
	config.LoadConfig()

	// Инициализация логгера
	logger.InitLogger()
	logger.Log.Info("Starting application")

	// Инициализация БД
	db, err := repository.NewPostgresDB()
	if err != nil {
		logger.Log.WithError(err).Fatal("Failed to initialize DB")
	}

	// Инициализация репозитория
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(userRepo)
	productHandler := handlers.NewProductHandler(productRepo)

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/api/register", authHandler.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/login", authHandler.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/api/products/category/{category_id}", productHandler.GetProductsByCategory).Methods("GET")
	r.HandleFunc("/api/favorites/add", authHandler.AddFavorite).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/favorites/remove", authHandler.RemoveFavorite).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/favorites", authHandler.GetFavorites).Methods("GET")
	r.HandleFunc("/api/basket/add", authHandler.AddToBasket).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/basket/remove", authHandler.RemoveFromBasket).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/basket", authHandler.GetBasket).Methods("GET")
	r.HandleFunc("/api/orders", authHandler.CreateOrder).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/basket/clear", authHandler.ClearBasket).Methods("POST", "OPTIONS")

	// Запуск сервера
	logger.Log.WithField("port", config.AppConfig.App.Port).Info("Server is running")
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.App.Port, enableCORS(r)))
}
