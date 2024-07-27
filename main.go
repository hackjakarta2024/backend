package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackjakarta2024/backend/config"
	"github.com/hackjakarta2024/backend/internal/handler"
	"github.com/hackjakarta2024/backend/internal/repository"
	"github.com/hackjakarta2024/backend/internal/service"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := config.NewDbPool()
	if err != nil {
		log.Fatal(err)
	}
	defer config.CloseDB(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/register", userHandler.Register)
	v1.Post("/login", userHandler.Login)

	app.Listen(":3000")
}
