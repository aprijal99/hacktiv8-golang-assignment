package repository

import "test/entity"

type ProductRepository interface {
	FindById(id int) *entity.Product
	FindAll() *[]entity.Product
}
