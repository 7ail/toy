package bootstrap

import (
	"github.com/7ail/toy/internal/repository"
	"github.com/7ail/toy/internal/user"
	"github.com/7ail/toy/pkg/zendesk"
)

func NewUserManager(id int) user.User {
	return user.New(id, repository.Manager(zendesk.Messenger()))
}