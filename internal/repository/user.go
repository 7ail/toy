package repository

import (
	"fmt"
	"errors"
)

type User struct {
	Id int
	Name string
}

func (m *manager) UserName(id int) (string, error) {
	user, err := m.user(id)
	if m.IsInternalServerError(err) {
		return "", fmt.Errorf("m.user(%v): %w", id, err)
	}

	if m.IsNotFoundError(err) {
		return m.fallbackUserName(id)
	}

	return user.Name, nil
}

func (m *manager) fallbackUserName(id int) (string, error) {
	name, err := m.fallback.UserName(id)
	if m.fallback.IsNotFoundError(err) {
		return "", fmt.Errorf("m.fallback.UserName(%v):%w%v", id, NotFoundError, err)
	}

	if e := m.fallback.IsRateLimitError(err); e != nil {
		return "", fmt.Errorf("m.fallback.UserName(%v):%w%v", id, newRateLimitError(e.RetryIn()), err)
	}

	if err != nil {
		return "", fmt.Errorf("m.fallback.UserName(%v):%w%v", id, InternalServerError, err)
	}

	return name, nil
}

func (m *manager) user(id int) (User, error) {
	switch id {
	case 1:
		return User{}, NotFoundError
	case 2:
		return User{}, NotFoundError
	default:
		return User{}, InternalServerError
	}
}

func (m *manager) IsNotFoundError(err error) bool {
	return errors.Is(err, NotFoundError)
}

func (m *manager) IsInternalServerError(err error) bool {
	return errors.Is(err, InternalServerError)
}

func (m *manager) IsRateLimitError(err error) *RateLimitError {
	var e *RateLimitError
	if errors.As(err, &e) {
		return e
	}

	return nil
}
