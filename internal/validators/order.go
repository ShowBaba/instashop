// validators/order.go
package validators

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"instashop/internal/dtos"
)

type OrderValidator struct {
	validate *validator.Validate
}

func NewOrderValidator() *OrderValidator {
	return &OrderValidator{validate: validator.New()}
}

func (v *OrderValidator) ValidatePlaceOrder(c *fiber.Ctx) error {
	var input dtos.PlaceOrderRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	if err := v.validate.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Validation failed",
			"errors":  err.(validator.ValidationErrors),
		})
	}

	c.Locals("input", input)
	return c.Next()
}
