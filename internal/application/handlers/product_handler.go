package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/clients"
	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/repository"
)

type ProductHandler interface {
	CreateProduct(ctx context.Context, req CreateProductRequest, user domain.UserInfo) (*domain.Product, error)
	GetAllProducts(ctx context.Context, user domain.UserInfo) ([]domain.Product, error)
	GetProductByID(ctx context.Context, user domain.UserInfo, productID string) (*domain.Product, error)
	GetUserInfo(token string) (*domain.UserInfo, error)
	UpdateProduct(ctx context.Context, id int64, req UpdateProductRequest) error
}

type productHandler struct {
	productRepository repository.ProductRepository
	userServiceClient clients.UserServiceClient
}

func NewProductHandler(productRepository repository.ProductRepository, userServiceClient clients.UserServiceClient) ProductHandler {
	return &productHandler{productRepository: productRepository, userServiceClient: userServiceClient}
}

type UpdateProductRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Quantity    *int    `json:"quantity,omitempty"`
	Image       *string `json:"image,omitempty"`
}

type CreateProductRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"categoryId"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

func (h *productHandler) GetUserInfo(token string) (*domain.UserInfo, error) {
	return h.userServiceClient.GetUserInfo(token)
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

func (h *productHandler) UpdateProduct(ctx context.Context, id int64, req UpdateProductRequest) error {
	product, err := h.productRepository.GetProductByID(ctx, id)

	if err != nil {
		return fmt.Errorf("could not get product by id: %v", err)
	}

	if product == nil {
		return fmt.Errorf("could not get product by id")
	}

	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return fmt.Errorf("yetkilendirme hatası: kullanıcı bilgisi bulunamadı")
	}

	if product.UserID != userID {
		return fmt.Errorf("yetkisiz erişim: bu ürünü güncelleme yetkiniz bulunmamaktadır")
	}

	if req.Title != nil {
		product.Title = *req.Title
	}

	if req.Description != nil {
		product.Description = *req.Description
	}

	if req.Quantity != nil {
		product.Quantity = *req.Quantity
	}

	if req.Image != nil {
		product.Image = *req.Image
	}

	product.UpdatedAt = time.Now()

	if err := h.productRepository.UpdateProduct(ctx, product); err != nil {
		return fmt.Errorf("could not update product: %v", err)
	}

	return nil
}

func (h *productHandler) GetAllProducts(ctx context.Context, user domain.UserInfo) ([]domain.Product, error) {
	products, err := h.productRepository.GetAllProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("ürünler alınırken hata oluştu: %v", err)
	}

	// Loglama ekleyelim
	log.Printf("GetAllProducts: Toplam %d ürün bulundu", len(products))

	return products, nil
}

func (h *productHandler) GetProductByID(ctx context.Context, user domain.UserInfo, productID string) (*domain.Product, error) {
	// String olan product ID'yi int64'e çevirelim
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("geçersiz ürün ID formatı: %v", err)
	}

	product, err := h.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ürün bilgisi alınırken hata oluştu: %v", err)
	}

	if product == nil {
		return nil, fmt.Errorf("ürün bulunamadı: ID=%d", id)
	}

	// Loglama ekleyelim
	log.Printf("GetProductByID: Ürün bulundu: ID=%d, Title=%s", id, product.Title)

	return product, nil
}
