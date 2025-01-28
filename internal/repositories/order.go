package repositories

import (
	"errors"

	"gorm.io/gorm"
	"instashop/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (o *OrderRepository) Create(order *models.Order) error {
	return o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for _, item := range order.Items {
			item.OrderID = order.ID
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (o *OrderRepository) FindByID(orderID uint) (*models.Order, bool, error) {
	var order models.Order
	if err := o.db.Preload("Items").First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &order, true, nil
}

func (o *OrderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := o.db.Preload("Items").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderRepository) UpdateStatus(orderID uint, status models.OrderStatus) error {
	return o.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (o *OrderRepository) Delete(orderID uint) error {
	return o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", orderID).Delete(&models.OrderItem{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Order{}, orderID).Error; err != nil {
			return err
		}
		return nil
	})
}
