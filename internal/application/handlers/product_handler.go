package handlers

import (
	"context"
	"time"

	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/clients"
	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/repository"
)

type ProductHandler interface {
	CreateProduct(ctx context.Context, req CreateProductRequest, user domain.UserInfo) (*domain.Product, error)
}

type productHandler struct {
	productRepository repository.ProductRepository
	userServiceClient clients.UserServiceClient
}

func NewProductHandler(productRepository repository.ProductRepository, userServiceClient clients.UserServiceClient) ProductHandler {
	return &productHandler{productRepository: productRepository, userServiceClient: userServiceClient}
}

type CreateProductRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"categoryId"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

func (h *productHandler) CreateProduct(ctx context.Context, req CreateProductRequest, user domain.UserInfo) (*domain.Product, error) {
	product := &domain.Product{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Image:       req.Image,
		UserID:      user.ID,
		Username:    user.Username,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := h.productRepository.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
