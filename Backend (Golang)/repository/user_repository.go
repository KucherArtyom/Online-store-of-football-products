package repository

import (
	"Go-Kurs/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO customers (name, surname, patronymic, telephone, login, password) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := r.db.QueryRow(
		query,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Telephone,
		user.Login,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsLoginExists(login string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM customers WHERE login = $1)`
	err := r.db.QueryRow(query, login).Scan(&exists)
	return exists, err
}

func (r *UserRepository) Authenticate(login, password string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, surname FROM customers WHERE login = $1 AND password = $2`
	err := r.db.QueryRow(query, login, password).Scan(
		&user.ID, // Убедитесь, что сканируется ID
		&user.Name,
		&user.Surname,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// user_repository.go
func (r *UserRepository) AddToFavorites(userID, productID int) error {
	// Проверка существования пользователя
	var userExists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM customers WHERE id = $1)", userID).Scan(&userExists)
	if err != nil || !userExists {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	// Проверка существования товара
	var productExists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)", productID).Scan(&productExists)
	if err != nil || !productExists {
		return fmt.Errorf("product with ID %d does not exist", productID)
	}

	// Транзакция
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("transaction begin error: %v", err)
	}
	defer tx.Rollback()

	// 1. Получаем или создаем запись в favourites
	var favID int
	err = tx.QueryRow(`
        WITH inserted AS (
            INSERT INTO favourites (customer_id)
            VALUES ($1)
            ON CONFLICT (customer_id) DO UPDATE SET customer_id = EXCLUDED.customer_id
            RETURNING id
        )
        SELECT id FROM inserted`, userID).Scan(&favID)

	if err != nil {
		return fmt.Errorf("favourites creation error: %v", err)
	}

	// 2. Добавляем связь продукта
	_, err = tx.Exec(`
        INSERT INTO product_favourites (product_id, favourites_id)
        VALUES ($1, $2)
        ON CONFLICT DO NOTHING`, productID, favID)

	if err != nil {
		return fmt.Errorf("product_favourites insertion error: %v", err)
	}

	return tx.Commit()
}

func (r *UserRepository) RemoveFromFavorites(userID, productID int) error {
	// Удаляем через JOIN для правильного определения связи
	result, err := r.db.Exec(`
        DELETE FROM product_favourites pf
        USING favourites f
        WHERE pf.favourites_id = f.id
        AND f.customer_id = $1
        AND pf.product_id = $2
    `, userID, productID)

	if err != nil {
		return fmt.Errorf("failed to remove favorite: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product %d not found in favorites for user %d", productID, userID)
	}

	return nil
}

func (r *UserRepository) GetFavorites(userID int) ([]models.Product, error) {
	query := `
        SELECT p.id, p.name, m.name as manufacturer, p.price, p.image_url, 
               p.description, p.category_id, p.quantity
        FROM products p
        JOIN manufacturers m ON p.manufacturer_id = m.id
        JOIN product_favourites pf ON p.id = pf.product_id
        JOIN favourites f ON pf.favourites_id = f.id
        WHERE f.customer_id = $1
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Manufacturer,
			&p.Price,
			&p.ImageURL,
			&p.Description,
			&p.CategoryID,
			&p.StockQuantity,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// user_repository.go
func (r *UserRepository) AddToBasket(userID, productID int) error {
	// Проверка существования пользователя
	var userExists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM customers WHERE id = $1)", userID).Scan(&userExists)
	if err != nil || !userExists {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	// Проверка существования товара
	var productExists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)", productID).Scan(&productExists)
	if err != nil || !productExists {
		return fmt.Errorf("product with ID %d does not exist", productID)
	}

	// Транзакция
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("transaction begin error: %v", err)
	}
	defer tx.Rollback()

	// 1. Получаем или создаем запись в baskets
	var basketID int
	err = tx.QueryRow(`
        WITH inserted AS (
            INSERT INTO baskets (customer_id)
            VALUES ($1)
            ON CONFLICT (customer_id) DO UPDATE SET customer_id = EXCLUDED.customer_id
            RETURNING id
        )
        SELECT id FROM inserted`, userID).Scan(&basketID)

	if err != nil {
		return fmt.Errorf("basket creation error: %v", err)
	}

	// 2. Добавляем связь продукта
	_, err = tx.Exec(`
        INSERT INTO product_baskets (product_id, baskets_id)
        VALUES ($1, $2)
        ON CONFLICT DO NOTHING`, productID, basketID)

	if err != nil {
		return fmt.Errorf("product_baskets insertion error: %v", err)
	}

	return tx.Commit()
}

func (r *UserRepository) RemoveFromBasket(userID, productID int) error {
	result, err := r.db.Exec(`
        DELETE FROM product_baskets pb
        USING baskets b
        WHERE pb.baskets_id = b.id
        AND b.customer_id = $1
        AND pb.product_id = $2
    `, userID, productID)

	if err != nil {
		return fmt.Errorf("failed to remove from basket: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product %d not found in basket for user %d", productID, userID)
	}

	return nil
}

func (r *UserRepository) GetBasket(userID int) ([]models.Product, error) {
	query := `
        SELECT p.id, p.name, m.name as manufacturer, p.price, p.image_url, 
               p.description, p.category_id, p.quantity
        FROM products p
        JOIN manufacturers m ON p.manufacturer_id = m.id
        JOIN product_baskets pb ON p.id = pb.product_id
        JOIN baskets b ON pb.baskets_id = b.id
        WHERE b.customer_id = $1
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Manufacturer,
			&p.Price,
			&p.ImageURL,
			&p.Description,
			&p.CategoryID,
			&p.StockQuantity,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *UserRepository) GetDB() *sql.DB {
	return r.db
}
