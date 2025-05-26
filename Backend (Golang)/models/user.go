package models

import "github.com/go-playground/validator/v10"

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name" validate:"required,min=2,max=20"`
	Surname    string `json:"surname" validate:"required,min=2,max=20"`
	Patronymic string `json:"patronymic" validate:"max=20"`
	Telephone  string `json:"telephone" validate:"required"`
	Login      string `json:"login" validate:"required,min=3,max=100"`
	Password   string `json:"password" validate:"required,min=6,max=100"`
}

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Manufacturer  string  `json:"manufacturer"`
	Price         float64 `json:"price"`
	ImageURL      string  `json:"image_url"`
	Description   string  `json:"description"`
	CategoryID    int     `json:"category_id"`
	StockQuantity int     `json:"stock_quantity"`
}

// LoginRequest представляет данные для авторизации пользователя.
type LoginRequest struct {
	Login    string `json:"login" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type Request struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type Request2 struct {
	UserID int `json:"user_id"`
}

type Address struct {
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	House     int    `json:"house"`
	Apartment int    `json:"apartment"`
}

type OrderRequest struct {
	CustomerID int     `json:"customer_id"`
	OrderPrice float64 `json:"order_price"`
	CardNumber string  `json:"card_number"`
	Products   []int   `json:"products"`
	Address    Address `json:"address"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
