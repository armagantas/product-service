package controllers

import (
	"context"
	"strings"

	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/handlers"
	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	GetAllProducts(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error
}

type productController struct {
	productHandler handlers.ProductHandler
}

func NewProductController(productHandler handlers.ProductHandler) *productController {
	return &productController{productHandler: productHandler}
}

func (c *productController) CreateProduct(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	token := parts[1]

	userInfo, err := c.productHandler.GetUserInfo(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	var req handlers.CreateProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	product, err := c.productHandler.CreateProduct(ctx.Context(), req, *userInfo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func (c *productController) UpdateProduct(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	token := parts[1]

	userInfo, err := c.productHandler.GetUserInfo(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	productID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	var req handlers.UpdateProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	customContext := context.WithValue(ctx.Context(), "userID", userInfo.ID)

	if err := c.productHandler.UpdateProduct(customContext, int64(productID), req); err != nil {
		if strings.Contains(err.Error(), "yetkisiz erişim") {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "yetkilendirme hatası") {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (c *productController) GetAllProducts(ctx *fiber.Ctx) error {
	var userInfo *domain.UserInfo
	authHeader := ctx.Get("Authorization")

	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			token := parts[1]
			userInfoResult, err := c.productHandler.GetUserInfo(token)
			if err == nil {
				userInfo = userInfoResult
			}
		}
	}

	if userInfo == nil {
		userInfo = &domain.UserInfo{}
	}

	categorySlug := ctx.Query("category")

	products, err := c.productHandler.GetAllProducts(ctx.Context(), *userInfo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ürünler yüklenirken hata oluştu: " + err.Error(),
		})
	}

	if categorySlug != "" && categorySlug != "all" {
		var filteredProducts []domain.Product
		for _, product := range products {
			if product.Category.Name == categorySlug {
				filteredProducts = append(filteredProducts, product)
			}
		}
		products = filteredProducts
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    products,
		"count":   len(products),
	})
}

func (c *productController) GetProductByID(ctx *fiber.Ctx) error {
	var userInfo *domain.UserInfo
	authHeader := ctx.Get("Authorization")

	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			token := parts[1]
			userInfoResult, err := c.productHandler.GetUserInfo(token)
			if err == nil {
				userInfo = userInfoResult
			}
		}
	}

	if userInfo == nil {
		userInfo = &domain.UserInfo{}
	}

	productID := ctx.Params("id")
	if productID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ürün ID'si gerekli",
		})
	}

	product, err := c.productHandler.GetProductByID(ctx.Context(), *userInfo, productID)
	if err != nil {
		if strings.Contains(err.Error(), "bulunamadı") {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ürün detayları yüklenirken hata oluştu: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    product,
	})
}
