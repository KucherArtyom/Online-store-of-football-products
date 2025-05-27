package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name" validate:"required,min=2,max=20"`
	Surname    string `json:"surname" validate:"required,min=2,max=20"`
	Patronymic string `json:"patronymic" validate:"max=20"`
	Telephone  string `json:"telephone" validate:"required"`
	Login      string `json:"login" validate:"required,min=3,max=100"`
	Password   string `json:"password" validate:"required,min=6,max=100"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type Claims struct {
	UserID int    `json:"user_id"`
	Login  string `json:"login"`
	jwt.RegisteredClaims
}

type Request struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type Request2 struct {
	UserID int `json:"user_id"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) GenerateJWT(secret string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: u.ID,
		Login:  u.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

type OrderRequest struct {
	CustomerID int     `json:"customer_id"`
	OrderPrice float64 `json:"order_price"`
	CardNumber string  `json:"card_number"`
	Products   []int   `json:"products"`
	Address    Address `json:"address"`
}

type Address struct {
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	House     int    `json:"house"`
	Apartment int    `json:"apartment"`
}
