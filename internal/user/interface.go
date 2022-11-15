package user

type source interface {
	IsNotFoundError(err error) (exist bool)
	IsRateLimitError(error) (retryIn int, err error)

	UserName(id int) (name string, err error)
}
