package zendesk

import (
	"fmt"
	"errors"
)

type User struct {
	Id int
	Name string
}

func (m *messenger) UserName(id int) (string, error) {
	user, err := m.user(id)
	if err != nil {
		return "", fmt.Errorf("m.user(%v): %w", id, err)
	}

	return user.Name, nil
}

// Mocked API call
func (m *messenger) user(id int) (User, error) {
	switch id {
	case 1:
		return User{}, NotFoundError
	case 2:
		return User{}, newRateLimitError(5)
	default:
		return User{}, InternalServerError
	}
}

func (m *messenger) IsNotFoundError(err error) bool {
	return errors.Is(err, NotFoundError)
}

func (m *messenger) IsInternalServerError(err error) bool {
	return errors.Is(err, InternalServerError)
}

func (m *messenger) IsRateLimitError(err error) *RateLimitError {
	var e *RateLimitError
	if errors.As(err, &e) {
		return e
	}

	return nil
}