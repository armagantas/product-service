package controllers

import (
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/handlers"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	AddCategory(c *fiber.Ctx) error
}

type categoryController struct {
	categoryHandler handlers.CategoryHandler
}

func NewCategoryController(categoryHandler handlers.CategoryHandler) *categoryController {
	return &categoryController{categoryHandler: categoryHandler}
}

func (c *categoryController) AddCategory(ctx *fiber.Ctx) error {
	var req handlers.CreateCategoryRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	category, err := c.categoryHandler.CreateCategory(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(category)

}
