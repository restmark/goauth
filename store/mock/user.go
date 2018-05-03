package mock

import (
	"github.com/restmark/goauth/helpers/params"
	"github.com/restmark/goauth/models"
)

func (m *mock) CreateUser(user *models.User) error {
	return nil
}

func (m *mock) FindUserById(id string) (*models.User, error) {
	user := &models.User{
		Id:        "123",
		Firstname: "maxence",
		Lastname:  "henneron",
		Email:     "maxence@restmark.co",
	}

	return user, nil
}

func (m *mock) FindUser(params params.M) (*models.User, error) {
	user := &models.User{
		Id:        "123",
		Firstname: "maxence",
		Lastname:  "henneron",
		Email:     "maxence@restmark.co",
	}

	return user, nil
}

func (m *mock) UpdateUser(user *models.User, params params.M) error {
	return nil
}
