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
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:5174",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	databaseConnection := postgresql.NewGormDB(&postgresql.Options{
		PgUsername: "postgres",
		PgPassword: "postgres",
		PgDbUrl:    "postgres://localhost:5432/product-service",
	})

	userServiceClient := clients.NewUserServiceClient("http://localhost:8001")

	productRepository := repository.NewProductRepository(databaseConnection)
	categoryRepository := repository.NewCategoryRepository(databaseConnection)

	productHandler := handlers.NewProductHandler(productRepository, userServiceClient)
	categoryHandler := handlers.NewCategoryHandler(categoryRepository)

	productController := controllers.NewProductController(productHandler)
	categoryController := controllers.NewCategoryController(categoryHandler)

	router.InitRouter(app, productController, categoryController)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World from Product Service!")
	})

	log.Println("Product service starting on port 8082...")
	log.Fatal(app.Listen(":8082"))
}
