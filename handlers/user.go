package handlers

import (
	"encoding/json"
	"github.com/restmark/goauth/models"
	"github.com/restmark/goauth/store"
)

func HandleUserCreated(store store.Store, message []byte) error {
	var user models.User

	err := json.Unmarshal(message, &user)
	if err != nil {
		return err
	}

	return store.CreateUser(&user)
}
