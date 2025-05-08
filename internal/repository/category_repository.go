package repository

import (
	"context"

	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}
