package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "github.com/guilimacode/lima-clm-api/docs"
	"github.com/guilimacode/lima-clm-api/shared/config"
	"github.com/guilimacode/lima-clm-api/shared/database"
)

// HealthCheck godoc
// @Summary      API health check
// @Description  Returns the current operational status of the API
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Lima CLM API is running",
	})
}

// @title           Lima CLM API
// @version         1.0
// @description     API for Lima CLM management.
// @host            localhost:8080
// @BasePath        /
func main() {

	config.LoadConfig()

	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "Lima CLM API",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health", HealthCheck)

	app.Get("/swagger/*", swagger.HandlerDefault)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	log.Printf("Server running on port %s...\n", port)

	<-quit
	log.Println("Starting Graceful Shutdown")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error while forcing shutdown: %v", err)
	}
}
