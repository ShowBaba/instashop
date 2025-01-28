package services

import (
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/repositories"
	"instashop/models"
)

type ProductClient interface {
	CreateProducts(inputs []dtos.CreateProductRequest) ([]models.Product, *common.RestErr)
	GetProduct(productID uint) (*models.Product, *common.RestErr)
	DeleteProduct(productID uint) *common.RestErr
	ListProducts() ([]models.Product, *common.RestErr)
	UpdateProduct(productID uint, input dtos.UpdateProductRequest) (*models.Product, *common.RestErr)
}

type ProductService struct {
	productRepo *repositories.ProductRepository
	restErr     *common.RestErr
}

func NewProductService(productRepo *repositories.ProductRepository,
	restErr *common.RestErr) ProductClient {
	return &ProductService{
		productRepo,
		restErr}
}

func (p *ProductService) CreateProducts(inputs []dtos.CreateProductRequest) ([]models.Product, *common.RestErr) {
	var products []models.Product
	var productNames []string

	for _, input := range inputs {
		productNames = append(productNames, input.Name)
	}

	existingProducts, err := p.productRepo.FetchByNames(productNames)
	if err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	existingProductMap := make(map[string]bool)
	for _, product := range existingProducts {
		existingProductMap[product.Name] = true
	}

	for _, input := range inputs {
		if _, exists := existingProductMap[input.Name]; exists {
			continue
		}
		product := models.Product{
			Name:        input.Name,
			Description: input.Description,
			Price:       input.Price,
			Stock:       input.Stock,
		}
		products = append(products, product)
	}

	if len(products) == 0 {
		return nil, p.restErr.BadRequest("No new products to add")
	}

	if err := p.productRepo.CreateMany(products); err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return products, nil
}

func (p *ProductService) GetProduct(productID uint) (*models.Product, *common.RestErr) {
	product, err := p.productRepo.FindByID(productID)
	if err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}
	if product == nil {
		return nil, p.restErr.BadRequest(common.ErrProductNotFound)
	}

	return product, nil
}

func (p *ProductService) UpdateProduct(productID uint, input dtos.UpdateProductRequest) (*models.Product, *common.RestErr) {
	product, err := p.productRepo.FindByID(productID)
	if err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}
	if product == nil {
		return nil, p.restErr.BadRequest(common.ErrProductNotFound)
	}

	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Price > 0 {
		product.Price = input.Price
	}
	if input.Stock >= 0 {
		product.Stock = input.Stock
	}

	if err := p.productRepo.Update(product); err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return product, nil
}

func (p *ProductService) DeleteProduct(productID uint) *common.RestErr {
	if err := p.productRepo.Delete(productID); err != nil {
		return p.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return nil
}

func (p *ProductService) ListProducts() ([]models.Product, *common.RestErr) {
	products, err := p.productRepo.ListAll()
	if err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return products, nil
}
