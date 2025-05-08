package main

import (
	"log"

	"github.com/armagantas/ecommerce-microservice/product-service/infrastructure/postgresql"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/clients"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/controllers"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/handlers"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/repository"
	"github.com/armagantas/ecommerce-microservice/product-service/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber uygulaması oluştur
	app := fiber.New()

	databaseConnection := postgresql.NewGormDB(&postgresql.Options{
		PgUsername: "test",
		PgPassword: "test",
		PgDbUrl:    "test",
	})

	userServiceClient := clients.NewUserServiceClient("http://localhost:8001")

	// Repository'leri oluştur
	productRepository := repository.NewProductRepository(databaseConnection)
	categoryRepository := repository.NewCategoryRepository(databaseConnection)

	// Handler'ları oluştur
	productHandler := handlers.NewProductHandler(productRepository, userServiceClient)
	categoryHandler := handlers.NewCategoryHandler(categoryRepository)

	// Controller'ları oluştur
	productController := controllers.NewProductController(productHandler)
	categoryController := controllers.NewCategoryController(categoryHandler)

	// Router'ı başlat
	router.InitRouter(app, productController, categoryController)

	// Ana route için Hello World endpoint'i
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World from Product Service!")
	})

	// 8082 portunda sunucuyu başlat
	log.Println("Product service starting on port 8082...")
	log.Fatal(app.Listen(":8082"))
}
