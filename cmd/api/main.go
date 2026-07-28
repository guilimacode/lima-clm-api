package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Lima CLM API",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Lima CLM API is running",
		})
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	log.Println("Servidor rodando na porta 8080...")

	<-quit
	log.Println("Iniciando Graceful Shutdown")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Erro ao forçar o desligamento: %v", err)
	}
}
