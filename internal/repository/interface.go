package repository

type source interface {
	UserName(id int) (name string, err error)

	IsNotFoundError(error) bool
	IsInternalServerError(error) bool
	IsRateLimitError(error) (retryIn int, err error)
}
