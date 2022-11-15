package repository

import (
	"github.com/7ail/toy/pkg/zendesk"
)

type source interface {
	UserName(id int) (name string, err error)

	IsNotFoundError(error) bool
	IsInternalServerError(error) bool
	IsRateLimitError(err error) *zendesk.RateLimitError
}
