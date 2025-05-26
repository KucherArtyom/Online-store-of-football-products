package repository

import (
	"Go-Kurs/models"
	"database/sql"
	"errors"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, m.name as manufacturer, p.price, p.image_url, 
		       p.description, p.category_id, p.quantity
		FROM products p
		JOIN manufacturers m ON p.manufacturer_id = m.id
	`
	rows, err := r.db.Query(query)
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

func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := `
		SELECT p.id, p.name, m.name as manufacturer, p.price, p.image_url,
		       p.description, p.category_id, p.quantity
		FROM products p
		JOIN manufacturers m ON p.manufacturer_id = m.id
		WHERE p.id = $1
		LIMIT 1
	`

	var product models.Product
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Manufacturer,
		&product.Price,
		&product.ImageURL,
		&product.Description,
		&product.CategoryID,
		&product.StockQuantity,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) GetProductsByCategory(categoryID int) ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, m.name as manufacturer, p.price, p.image_url,
		       p.description, p.category_id, p.quantity
		FROM products p
		JOIN manufacturers m ON p.manufacturer_id = m.id
		WHERE p.category_id = $1
	`

	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to query products by category: %w", err)
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
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found in category %d", categoryID)
	}

	return products, nil
}
