package router

import (
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(router fiber.Router, productController controllers.ProductController, categoryController controllers.CategoryController) {
	productRouter := router.Group("/products")
	categoryRouter := router.Group("/categories")

	productRouter.Post("/", productController.CreateProduct)
	productRouter.Put("/:id", productController.UpdateProduct)
	productRouter.Get("/", productController.GetAllProducts)
	productRouter.Get("/:id", productController.GetProductByID)
	productRouter.Get("/:id/owner", productController.GetProductOwner)
	categoryRouter.Post("/", categoryController.AddCategory)
}
