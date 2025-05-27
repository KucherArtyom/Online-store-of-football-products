package main

import (
	"Go-Kurs/config"
	"Go-Kurs/handlers"
	"Go-Kurs/logger"
	"Go-Kurs/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	config.LoadConfig()
	logger.InitLogger()
	logger.Log.Info("Starting application")

	db, err := repository.NewPostgresDB()
	if err != nil {
		logger.Log.WithError(err).Fatal("Failed to initialize DB")
	}

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	authHandler := handlers.NewAuthHandler(userRepo)
	productHandler := handlers.NewProductHandler(productRepo)

	r := mux.NewRouter()

	r.HandleFunc("/api/register", authHandler.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/login", authHandler.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/api/products/category/{category_id}", productHandler.GetProductsByCategory).Methods("GET")

	authRouter := r.PathPrefix("/api").Subrouter()
	authRouter.Use(authHandler.AuthMiddleware)

	authRouter.HandleFunc("/favorites/add", authHandler.AddFavorite).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/favorites/remove", authHandler.RemoveFavorite).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/favorites", authHandler.GetFavorites).Methods("GET")
	authRouter.HandleFunc("/basket/add", authHandler.AddToBasket).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/basket/remove", authHandler.RemoveFromBasket).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/basket", authHandler.GetBasket).Methods("GET")
	authRouter.HandleFunc("/orders", authHandler.CreateOrder).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/basket/clear", authHandler.ClearBasket).Methods("POST", "OPTIONS")

	logger.Log.WithField("port", config.AppConfig.App.Port).Info("Server is running")
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.App.Port, enableCORS(r)))
}
