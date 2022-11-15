package zendesk

import (
	"fmt"
)

var NotFoundError = fmt.Errorf("not found error")
var InternalServerError = fmt.Errorf("internal server error")

type RateLimitError struct {
	retryIn int
	Err     error
}

func newRateLimitError(retryIn int) *RateLimitError {
	return &RateLimitError{
		retryIn: retryIn,
		Err:     fmt.Errorf("rate limit error"),
	}
}

func (e *RateLimitError) Error() string {
	return e.Err.Error()
}
