package store

import (
	"github.com/restmark/goauth/models"
	"github.com/restmark/goauth/helpers/params"
)

type Store interface {
	CreateUser(*models.User) error
	FindUserById(string) (*models.User, error)
	FindUser(params.M) (*models.User, error)
	UpdateUser(*models.User, params.M) error
}
