package repository

import (
	"github.com/denny-idtrust/go_say_hello/v2/entity"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
