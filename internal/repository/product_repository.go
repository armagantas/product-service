package repository

import (
	"context"
	"errors"
	"time"

	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	GetProductByID(ctx context.Context, productID int64) (*domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository { return &productRepository{db: db} }

func (r *productRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.WithContext(ctx).Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetProductByID(ctx context.Context, productID int64) (*domain.Product, error) {
	var product domain.Product

	err := r.db.WithContext(ctx).Preload("Category").First(&product, productID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product *domain.Product) error {
	product.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *productRepository) GetUserByProductId(ctx context.Context, productID int64) error {
	return nil
}
