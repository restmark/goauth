package models

import (
	"strings"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/asaskevich/govalidator"
	"github.com/restmark/goauth/helpers"
)

type User struct {
	Id            string       `json:"id" bson:"_id,omitempty" valid:"-"`
	Firstname     string       `json:"first_name" bson:"first_name" valid:"required"`
	Lastname      string       `json:"last_name" bson:"last_name" valid:"required"`
	Password      string       `json:"password" bson:"password" valid:"required"`
	Email         string       `json:"email" bson:"email" valid:"email,required"`
}

type SanitizedUser struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
}


func (user *User) BeforeCreate() error {
	user.Email = strings.ToLower(user.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, "encryption_failed", "Failed to generate the crypted password")
	}
	user.Password = string(hashedPassword)

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, "input_not_valid", err.Error())
	}

	return nil
}

func (user *User) Sanitize() SanitizedUser {
	return SanitizedUser{user.Id, user.Firstname, user.Lastname, user.Email}
}

const UsersCollection = "users"
