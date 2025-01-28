package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/services"
	"instashop/internal/utils"
)

type OrderHandler struct {
	orderSvc services.OrderClient
	restErr  *common.RestErr
}

func NewOrderHandler(orderSvc services.OrderClient,
	restErr *common.RestErr,
) *OrderHandler {
	return &OrderHandler{
		orderSvc,
		restErr,
	}
}

func (o *OrderHandler) PlaceOrder(c *fiber.Ctx) error {
	var input dtos.PlaceOrderRequest
	i := c.Locals("input")
	input, ok := i.(dtos.PlaceOrderRequest)
	if !ok {
		log.Error(fmt.Errorf("cannot convert validated data to PlaceOrderRequest"))
		err := o.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}
	userID, err := utils.GetAuthUserIdFromContext(c)
	if err != nil {
		log.Error(zap.Error(err))
		err := o.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}
	order, srvErr := o.orderSvc.PlaceOrder(input, userID)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Order placed successfully",
		"data":    order,
	})
}

func (o *OrderHandler) ListOrders(c *fiber.Ctx) error {
	userID, err := utils.GetAuthUserIdFromContext(c)
	if err != nil {
		log.Error(zap.Error(err))
		err := o.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}
	orders, srvErr := o.orderSvc.ListOrders(userID)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Orders retrieved successfully",
		"data":    orders,
	})
}

func (o *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	orderID, err := c.ParamsInt("orderID")
	if err != nil {
		err := o.restErr.BadRequest(common.ErrInvalidOrder)
		return c.Status(err.StatusCode).JSON(err)

	}
	userID, err := utils.GetAuthUserIdFromContext(c)
	if err != nil {
		log.Error(zap.Error(err))
		err := o.restErr.ServerError(common.ErrSomethingWentWrong)
		return c.Status(err.StatusCode).JSON(err)
	}

	srvErr := o.orderSvc.CancelOrder(uint(orderID), userID)
	if srvErr != nil {
		return c.Status(srvErr.StatusCode).JSON(srvErr)
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Order canceled successfully",
	})
}
