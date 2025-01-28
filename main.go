package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db "instashop/database"
	"instashop/internal/utils"
	"instashop/router"
)

func main() {
	err := db.ConnectToPgDB(
		utils.GetConfig().DbHost,
		utils.GetConfig().DbUser,
		utils.GetConfig().DbPassword,
		utils.GetConfig().DbName,
		utils.GetConfig().DbPort,
	)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	loggerSettings := logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})
	app.Use(loggerSettings)

	router.Routes(app, db.Client)

	db.Migrate(db.Client)

	if err := db.StartSeeder(db.Client); err != nil {
		fmt.Printf("Failed to seed data: %v\n", err)
	}

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := utils.GetConfig().Port

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}

		close(idleConnsClosed)
	}()

	log.Printf("Starting server on port: %s", port)
	if err := app.Listen(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Listen: %v", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped gracefully")
}
