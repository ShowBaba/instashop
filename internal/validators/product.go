package validators

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"instashop/internal/dtos"
)

type ProductValidator struct {
	validate *validator.Validate
}

func NewProductValidator() *ProductValidator {
	return &ProductValidator{validate: validator.New()}
}

func (v *ProductValidator) ValidateCreateProduct(c *fiber.Ctx) error {
	var input dtos.CreateProductRequest
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

func (v *ProductValidator) ValidateUpdateProduct(c *fiber.Ctx) error {
	var input dtos.UpdateProductRequest
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
