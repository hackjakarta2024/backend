package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackjakarta2024/backend/config"
	"github.com/hackjakarta2024/backend/internal/handler"
	"github.com/hackjakarta2024/backend/internal/middleware"
	"github.com/hackjakarta2024/backend/internal/repository"
	"github.com/hackjakarta2024/backend/internal/service"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file")
	}

	db, err := config.NewDbPool()
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}
	defer config.CloseDB(db)

	bq, err := config.NewBigQuery()
	if err != nil {
		logger.Fatal("Error connecting to BigQuery", zap.Error(err))
	}
	defer config.CloseBigQuery(bq)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	promoRepository := repository.NewPromoRepository(db)
	foodRepository := repository.NewFoodRepository(db)
	restaurantRepository := repository.NewRestaurantRepository(db)
	fypRepository := repository.NewFypRepository(db, bq, logger)
	fypService := service.NewFypService(fypRepository, promoRepository, foodRepository, restaurantRepository, logger)
	fypHandler := handler.NewFypHandler(fypService, logger)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/register", userHandler.Register)
	v1.Post("/login", userHandler.Login)

	fyp := v1.Group("/fyp")
	fyp.Get("/food", middleware.JWTMiddleware(), fypHandler.GetFyp)

	app.Listen(":3000")
}
