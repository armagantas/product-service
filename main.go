package main

import (
	"log"

	"github.com/armagantas/ecommerce-microservice/product-service/infrastructure"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Fiber uygulaması oluştur
	app := fiber.New()

	db, err := infrastructure.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Database connected successfully", db)

	// Ana route için Hello World endpoint'i
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World from Product Service!")
	})

	// 8082 portunda sunucuyu başlat
	log.Println("Product service starting on port 8082...")
	log.Fatal(app.Listen(":8082"))
} 