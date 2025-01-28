package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/services"
)

type AuthHandler struct {
	authSvc services.AuthClient
	restErr *common.RestErr
}

func NewAuthHandler(authSvc services.AuthClient, restErr *common.RestErr) *AuthHandler {
	return &AuthHandler{
		authSvc,
		restErr,
	}
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	var input dtos.LoginDTO

	i := c.Locals("input")
	input, ok := i.(dtos.LoginDTO)
	if !ok {
		log.Error(fmt.Errorf("cannot convert validated data to LoginDTO"))
		err := a.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}
	resp, srvErr := a.authSvc.Login(input)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	c.Status(200)
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "login successful",
		"data":    resp,
	})
}

func (a *AuthHandler) Signup(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")

	var input dtos.SignUpDTO

	i := c.Locals("input")
	input, ok := i.(dtos.SignUpDTO)
	if !ok {
		log.Error(fmt.Errorf("cannot convert validated data to SignUpDTO"))
		err := a.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	resp, err := a.authSvc.Signup(input)
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	c.Status(201)
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "signup successful",
		"data":    resp,
	})
}
