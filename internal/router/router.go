package router

import (
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(router fiber.Router, productController controllers.ProductController, categoryController controllers.CategoryController) {
	productRouter := router.Group("/products")
	categoryRouter := router.Group("/categories")

	productRouter.Post("/", productController.CreateProduct)
	categoryRouter.Post("/", categoryController.AddCategory)
}
