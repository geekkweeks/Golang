package services

import (
	"learn-go-unit-test/entity"
	"learn-go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}


func TestCategoryService_GetNotFound(t *testing.T){
	// simulasi mock
	// kita bisa membuat object dengan Mock berserta return data yg diharapkan
	categoryRepository.Mock.On("FindById","1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}


func TestCategoryService_GetFound(t *testing.T){
	category := entity.Category{
		Id: "1",
		Name: "Laptop",
	}
	// simulasi mock
	// kita bisa membuat object dengan Mock berserta return data yg diharapkan
	categoryRepository.Mock.On("FindById","2").Return(nil)

	result, err := categoryService.Get("2")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}
