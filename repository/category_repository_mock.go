package repository

import (
	"github.com/denny-idtrust/go_say_hello/v2/entity"
	"github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock is a mock of CategoryRepository interface
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (c *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := c.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		return args.Get(0).(*entity.Category)
	}
}
