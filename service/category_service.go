package service

import (
	"errors"
	"github.com/denny-idtrust/go_say_hello/v2/entity"
	"github.com/denny-idtrust/go_say_hello/v2/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
