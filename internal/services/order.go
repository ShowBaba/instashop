package services

import (
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/repositories"
	"instashop/models"
)

type OrderClient interface {
	PlaceOrder(input dtos.PlaceOrderRequest, userID uint) (*models.Order, *common.RestErr)
	ListOrders(userID uint) ([]dtos.OrderResponse, *common.RestErr)
	CancelOrder(userID, orderID uint) *common.RestErr
}

type OrderService struct {
	orderRepo   *repositories.OrderRepository
	productRepo *repositories.ProductRepository
	restErr     *common.RestErr
}

func NewOrderService(
	orderRepo *repositories.OrderRepository,
	productRepo *repositories.ProductRepository,
	restErr *common.RestErr,
) OrderClient {
	return &OrderService{
		orderRepo,
		productRepo,
		restErr,
	}
}

func (o *OrderService) PlaceOrder(input dtos.PlaceOrderRequest, userID uint) (*models.Order, *common.RestErr) {
	var totalPrice float64
	var orderItems []models.OrderItem

	for _, item := range input.Items {
		product, err := o.productRepo.FindByID(item.ProductID)
		if err != nil {
			return nil, o.restErr.ServerError(common.ErrSomethingWentWrong)
		}
		if product == nil {
			return nil, o.restErr.BadRequest(common.ErrProductNotFound)
		}

		if product.Stock < item.Quantity {
			return nil, o.restErr.BadRequest(common.ErrInsufficientStock)
		}

		itemPrice := product.Price * float64(item.Quantity)
		totalPrice += itemPrice

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		if err := o.productRepo.DecreaseStock(item.ProductID, item.Quantity); err != nil {
			return nil, o.restErr.ServerError(common.ErrSomethingWentWrong)
		}
	}

	order := models.Order{
		UserID:     userID,
		Status:     models.OrderStatusPending,
		TotalPrice: totalPrice,
		Items:      orderItems,
	}

	if err := o.orderRepo.Create(&order); err != nil {
		return nil, o.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return &order, nil
}

func (o *OrderService) ListOrders(userID uint) ([]dtos.OrderResponse, *common.RestErr) {
	orders, err := o.orderRepo.FindByUserID(userID)
	if err != nil {
		return nil, o.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	var orderResponses []dtos.OrderResponse
	for _, order := range orders {
		orderResponse := dtos.OrderResponse{
			ID:         order.ID,
			UserID:     order.UserID,
			Status:     string(order.Status),
			TotalPrice: order.TotalPrice,
			CreatedAt:  order.CreatedAt,
			Items:      make([]dtos.OrderItemDetail, 0),
		}

		for _, item := range order.Items {
			orderResponse.Items = append(orderResponse.Items, dtos.OrderItemDetail{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			})
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func (o *OrderService) CancelOrder(userID, orderID uint) *common.RestErr {

	order, exists, err := o.orderRepo.FindByID(uint(orderID))
	if err != nil {
		return o.restErr.ServerError(common.ErrSomethingWentWrong)
	}
	if !exists {
		return o.restErr.BadRequest(common.ErrOrderNotFound)
	}

	if order.UserID != userID {
		return o.restErr.BadRequest(common.ErrUnauthorized)
	}

	if order.Status != models.OrderStatusPending {
		return o.restErr.BadRequest(common.ErrCanOnlyCancelPendingOrder)
	}

	if err := o.orderRepo.UpdateStatus(uint(orderID), models.OrderStatusCancelled); err != nil {
		return o.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return nil
}
