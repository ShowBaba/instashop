package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"instashop/internal/common"
	"instashop/internal/services"
	"instashop/internal/utils"
)

type UserHandler struct {
	userSvc services.UserClient
	restErr *common.RestErr
}

func NewUserHandler(
	userSvc services.UserClient,
	restErr *common.RestErr,
) *UserHandler {
	return &UserHandler{
		userSvc,
		restErr,
	}
}

func (u *UserHandler) GetUserDetails(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	userId, err := utils.GetAuthUserIdFromContext(c)
	if err != nil {
		log.Error(zap.Error(err))
		err := u.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	resp, srvErr := u.userSvc.GetUserDetails(userId)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	c.Status(200)
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "fetched user data successful",
		"data":    resp,
	})
}
