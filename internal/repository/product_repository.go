package repository

import (
	"context"

	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository { return &productRepository{db: db} }

func (r *productRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}
