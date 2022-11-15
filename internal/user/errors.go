package user

import (
	"fmt"
)

var NotFoundError = fmt.Errorf(" ")
var InternalServerError = fmt.Errorf(" ")

type RateLimitError struct {
	retryIn int
	Err     error
}

func newRateLimitError(retryIn int) *RateLimitError {
	return &RateLimitError{
		retryIn: retryIn,
		Err:     fmt.Errorf(" "),
	}
}

func (e *RateLimitError) Error() string {
	return e.Err.Error()
}
