package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"instashop/internal/common"
	"instashop/internal/handlers"
	"instashop/internal/middleware"
	"instashop/internal/repositories"
	"instashop/internal/services"
)

func RegisterUserRoutes(router fiber.Router, db *gorm.DB) {
	restErr := common.NewRestErr()
	authMiddleware := middleware.NewAuthMiddleware(restErr)
	userRepo := repositories.NewUserRepository(db)

	userSvc := services.NewUserService(userRepo, restErr)
	handler := handlers.NewUserHandler(userSvc, restErr)

	userRouter := router.Group("user")
	userRouter.Get("/get-details", authMiddleware.ValidateAuthHeaderToken, handler.GetUserDetails)
}
