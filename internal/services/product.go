package services

import (
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/repositories"
	"instashop/models"
)

type ProductClient interface {
	CreateProduct(input dtos.CreateProductRequest) (*models.Product, *common.RestErr)
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

func (p *ProductService) CreateProduct(input dtos.CreateProductRequest) (*models.Product, *common.RestErr) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	if err := p.productRepo.Create(&product); err != nil {
		return nil, p.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return &product, nil
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
