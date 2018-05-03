package handlers

import (
	"github.com/restmark/goauth/store/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var storeMock = mock.New()

func TestHandleUserCreated(t *testing.T) {
	var userJson = []byte(`{
	  "first_name": "maxence",
	  "last_name": "henneron",
	  "password":"test",
	  "email":"maxence@restmark.co"
	}`)

	err := HandleUserCreated(storeMock, userJson)
	assert.Nil(t, err)
}

func TestHandleUserCreatedWithError(t *testing.T) {
	var userJson = []byte(`
    	invalid json
	`)

	err := HandleUserCreated(storeMock, userJson)
	assert.NotNil(t, err)
}
