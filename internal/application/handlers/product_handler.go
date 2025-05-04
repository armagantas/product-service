package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/clients"
	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/repository"
)

type ProductHandler interface {
	CreateProduct(ctx context.Context, w http.ResponseWriter, r *http.Request)
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


func (h *productHandler) CreateProduct(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := extractBearerToken(r.Header.Get("Authorization"))
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userInfo, err := h.userServiceClient.GetUserInfo(token)
	if err != nil {
		http.Error(w, "User information cannot get from client", http.StatusUnauthorized)
		return
	}

	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product := &domain.Product{
		Title: req.Title,
		Description: req.Description,
		CategoryID: req.CategoryID,
		Quantity: req.Quantity,
		Price: req.Price,
		Image: req.Image,
		UserID: userInfo.ID,
		Username: userInfo.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.productRepository.CreateProduct(ctx, product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
	
	
}

func extractBearerToken(authHeader string) string {
	const prefix = "Bearer "

	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		return ""
	}

	return authHeader[len(prefix):]
}
