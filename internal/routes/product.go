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

func RegisterProductRoutes(router fiber.Router, db *gorm.DB) {
	restErr := common.NewRestErr()
	authMiddleware := middleware.NewAuthMiddleware(restErr)
	productRepo := repositories.NewProductRepository(db)
	productSvc := services.NewProductService(productRepo, restErr)
	productValidator := validators.NewProductValidator()
	productHandler := handlers.NewProductHandler(productSvc, restErr)

	productRouter := router.Group("product")
	productRouter.Use(authMiddleware.ValidateAuthHeaderToken)
	productRouter.Use(middleware.AdminOnly)

	productRouter.Post("/create-product", productValidator.ValidateCreateProduct, productHandler.CreateProduct)
	productRouter.Get("/list-products", productHandler.ListProducts)
	productRouter.Get("/:productID", productHandler.GetProduct)
	productRouter.Patch("/:productID", productValidator.ValidateUpdateProduct, productHandler.UpdateProduct)
	productRouter.Delete("/:productID", productHandler.DeleteProduct)
}
