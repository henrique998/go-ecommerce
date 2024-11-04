package repositories

import (
	"context"
	"net/http"
	"time"

	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type pgProductsRepository struct {
	db *pgx.Conn
}

func NewPGProductsRepo(db *pgx.Conn) contracts.ProductsRepository {
	return &pgProductsRepository{db: db}
}

func (r *pgProductsRepository) FindMany(page, size int) ([]models.Product, error) {
	offset := (page - 1) * size
	query := `SELECT id, name, description, price, stock_quantity, image_url, external_id, created_at, updated_at from products LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(context.Background(), query, size, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var id, name, description, imageUrl, externalID string
	var price int64
	var stockQty int
	var createdAt time.Time
	var updatedAt *time.Time

	var products []models.Product

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &stockQty, &imageUrl, &externalID, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		p := models.NewExistingProduct(id, name, description, price, stockQty, imageUrl, externalID, createdAt, updatedAt)
		products = append(products, p)
	}

	return products, nil
}

func (r *pgProductsRepository) Create(product models.Product) errors.AppErr {
	query := `INSERT INTO products (id, name, description, price, stock_quantity, image_url, external_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.GetID(),
		product.GetName(),
		product.GetDescription(),
		product.GetPrice(),
		product.GetStock(),
		product.GetImageUrl(),
		product.GetExternalID(),
		product.GetCreatedAt(),
		product.GetUpdatedAt(),
	)
	if err != nil {
		return errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	return nil
}
