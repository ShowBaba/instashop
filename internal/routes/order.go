package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"instashop/internal/common"
	"instashop/internal/handlers"
	"instashop/internal/middleware"
	"instashop/internal/repositories"
	"instashop/internal/services"
	"instashop/internal/validators"
)

func RegisterOrderRoutes(router fiber.Router, db *gorm.DB) {
	restErr := common.NewRestErr()
	authMiddleware := middleware.NewAuthMiddleware(restErr)
	orderRepo := repositories.NewOrderRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderSvc := services.NewOrderService(orderRepo, productRepo, restErr)
	orderValidator := validators.NewOrderValidator()
	orderHandler := handlers.NewOrderHandler(orderSvc, restErr)

	orderRouter := router.Group("order")
	orderRouter.Use(authMiddleware.ValidateAuthHeaderToken)

	orderRouter.Post("/place-order", orderValidator.ValidatePlaceOrder, orderHandler.PlaceOrder)
	orderRouter.Get("/list-order", orderHandler.ListOrders)
	orderRouter.Patch("/:orderID/cancel", orderHandler.CancelOrder)
}
