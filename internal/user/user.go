package user

import (
	"errors"
	"fmt"
)

type User struct {
	id         int
	repository source
}

func New(id int, r source) User {
	return User{
		id:         id,
		repository: r,
	}
}

func (u *User) Name() (string, error) {
	name, err := u.repository.UserName(u.id)
	if u.repository.IsNotFoundError(err) {
		return "", fmt.Errorf("u.repository.UserName(%v):%w%v", u.id, NotFoundError, err)
	}

	if retryIn, e := u.repository.IsRateLimitError(err); e == nil {
		return "", fmt.Errorf("m.repository.UserName(%v):%w%v", u.id, newRateLimitError(retryIn), err)
	}

	if err != nil {
		return "", fmt.Errorf("u.repository.UserName(%v):%w%v", u.id, InternalServerError, err)
	}

	return name, nil
}

func (u *User) IsNotFoundError(err error) bool {
	return errors.Is(err, NotFoundError)
}

func (u *User) IsInternalServerError(err error) bool {
	return errors.Is(err, InternalServerError)
}

func (u *User) IsRateLimitError(err error) (int, error) {
	var e *RateLimitError
	if errors.As(err, &e) {
		return e.retryIn, nil
	}

	return 0, err
}
