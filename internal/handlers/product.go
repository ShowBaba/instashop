package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/services"
)

type ProductHandler struct {
	productSvc services.ProductClient
	restErr    *common.RestErr
}

func NewProductHandler(productSvc services.ProductClient,
	restErr *common.RestErr) *ProductHandler {
	return &ProductHandler{productSvc, restErr}
}

func (p *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var input []dtos.CreateProductRequest
	if err := c.BodyParser(&input); err != nil {
		log.Error(zap.Error(err))
		err := p.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	product, srvErr := p.productSvc.CreateProducts(input)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Product created successfully",
		"data":    product,
	})
}

func (p *ProductHandler) GetProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("productID")
	if err != nil {
		log.Error(zap.Error(err))
		err := p.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	product, srvErr := p.productSvc.GetProduct(uint(productID))
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

func (p *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("productID")
	if err != nil {
		log.Error(zap.Error(err))
		err := p.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	var input dtos.UpdateProductRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	product, srvErr := p.productSvc.UpdateProduct(uint(productID), input)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product updated successfully",
		"data":    product,
	})
}

func (p *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("productID")
	if err != nil {
		log.Error(zap.Error(err))
		err := p.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	if srvErr := p.productSvc.DeleteProduct(uint(productID)); srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product deleted successfully",
	})
}

func (p *ProductHandler) ListProducts(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid page parameter",
		})
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid pageSize parameter",
		})
	}

	products, totalCount, srvErr := p.productSvc.ListProducts(page, pageSize)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	totalPages := (int(totalCount) + pageSize - 1) / pageSize
	nextPage := page + 1
	if page >= totalPages {
		nextPage = 0
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Products retrieved successfully",
		"data":    products,
		"pagination": fiber.Map{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalPages":  totalPages,
			"totalCount":  totalCount,
			"nextPage":    nextPage,
		},
	})
}
