package controllers

import (
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/handlers"
	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
}

type productController struct {
	productHandler handlers.ProductHandler
}

func NewProductController(productHandler handlers.ProductHandler) *productController {
	return &productController{productHandler: productHandler}
}

func (c *productController) CreateProduct(ctx *fiber.Ctx) error {
	var req handlers.CreateProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Temporary hardcoded user info
	userInfo := domain.UserInfo{
		ID:       "1",
		Username: "admin",
	}

	product, err := c.productHandler.CreateProduct(ctx.Context(), req, userInfo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(product)
}
