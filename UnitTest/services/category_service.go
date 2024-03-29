package services

import (
	"errors"
	"learn-go-unit-test/entity"
	"learn-go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string)(*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category is not found")
	} else{
		return category, nil
	}
}