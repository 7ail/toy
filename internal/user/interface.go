package user

import (
	"github.com/7ail/toy/internal/repository"
)

type source interface {
	IsNotFoundError(err error) (exist bool)
	IsRateLimitError(err error) *repository.RateLimitError

	UserName(id int) (name string, err error)
}
