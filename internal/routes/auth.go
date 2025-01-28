package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"instashop/internal/common"
	"instashop/internal/handlers"
	"instashop/internal/repositories"
	"instashop/internal/services"
	"instashop/internal/validators"
)

func RegisterAuthRoutes(router fiber.Router, db *gorm.DB) {
	restErr := common.NewRestErr()
	userRepo := repositories.NewUserRepository(db)
	authSvc := services.NewAuthService(userRepo, restErr)
	validator := validators.NewAuthValidator(userRepo, restErr)
	handler := handlers.NewAuthHandler(authSvc, restErr)

	userRouter := router.Group("auth")
	userRouter.Post("/login", validator.ValidateLogin, handler.Login)
	userRouter.Post("/signup", validator.ValidateSignup, handler.Signup)
}
