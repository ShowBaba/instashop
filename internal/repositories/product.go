package repositories

import (
	"errors"

	"gorm.io/gorm"
	"instashop/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) Create(product *models.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductRepository) CreateMany(products []models.Product) error {
	return p.db.Create(&products).Error
}
func (p *ProductRepository) FindByID(productID uint) (*models.Product, error) {
	var product models.Product
	if err := p.db.First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) FetchByNames(names []string) ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Where("name IN ?", names).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) Update(product *models.Product) error {
	return p.db.Save(product).Error
}

func (p *ProductRepository) Delete(productID uint) error {
	return p.db.Delete(&models.Product{}, productID).Error
}

func (p *ProductRepository) ListAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) DecreaseStock(productID uint, quantity int) error {
	return p.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Update("stock", gorm.Expr("stock - ?", quantity)).
		Error
}
