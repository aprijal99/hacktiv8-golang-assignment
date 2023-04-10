package service

import (
	"test/entity"
	"test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (productService ProductService) GetOneProduct(id int) *entity.Product {
	product := productService.Repository.FindById(id)
	if product == nil {
		return nil
	}

	return product
}

func (ProductService ProductService) GetAllProducts() *[]entity.Product {
	products := ProductService.Repository.FindAll()
	if products == nil {
		return nil
	}

	return products
}
