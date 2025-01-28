package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"instashop/internal/common"
	"instashop/internal/utils"
)

type AuthMiddleware struct {
	restErr *common.RestErr
}

func NewAuthMiddleware(
	restErr *common.RestErr,
) *AuthMiddleware {
	return &AuthMiddleware{
		restErr,
	}
}

func (a *AuthMiddleware) ValidateAuthHeaderToken(c *fiber.Ctx) error {
	tokenInHeader := c.Get("Authorization")
	if tokenInHeader == "" {
		return c.Status(http.StatusBadRequest).JSON(a.restErr.ServerError(common.ErrMissingAuthTokenInHeader))
	}
	token := strings.Split(tokenInHeader, " ")[1]
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(a.restErr.ServerError(common.ErrMissingAuthTokenInHeader))
	}
	claim, err := utils.ValidateAuthToken(token, utils.GetConfig().JWTSecretKey)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(a.restErr.ServerError(common.ErrInvalidAuthToken))
	}

	c.Set("email", claim.Email)
	c.Set("role", claim.Role)
	c.Set("id", strconv.FormatUint(uint64(claim.ID), 10))
	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	userID, err := utils.GetAuthUserIdFromContext(c)
	if err != nil {
		log.Error(zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(common.ErrSomethingWentWrong)
	}
	userRole := c.GetRespHeader("role")
	if userRole != "admin" {
		c.Status(200)
		return c.JSON(&fiber.Map{
			"success": true,
			"message": common.ErrUnauthorized,
		})
	}

	fmt.Println("userID:", userID)

	return c.Next()
}
