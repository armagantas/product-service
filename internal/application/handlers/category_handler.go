package handlers

import (
	"context"
	"time"

	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/repository"
)

type CategoryHandler interface {
	CreateCategory(ctx context.Context, req CreateCategoryRequest) (*domain.Category, error)
}

type categoryHandler struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryHandler(categoryRepository repository.CategoryRepository) CategoryHandler {
	return &categoryHandler{categoryRepository: categoryRepository}
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (h *categoryHandler) CreateCategory(ctx context.Context, req CreateCategoryRequest) (*domain.Category, error) {
	category := &domain.Category{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.categoryRepository.CreateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
