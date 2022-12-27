package repository

import "learn-go-unit-test/entity"

type CategoryRepository interface {
	// return pointer category
	// contract
	FindById(id string) *entity.Category
}