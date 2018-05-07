package controllers

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/config"
	"github.com/restmark/goauth/helpers"
	"github.com/restmark/goauth/helpers/params"
	"github.com/restmark/goauth/models"
	"github.com/restmark/goauth/store"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthController struct {
}

func NewAuthController() AuthController {
	return AuthController{}
}

func (ac AuthController) Authentication(c *gin.Context) {

	//Read base64 private key
	encodedKey := []byte(config.GetString(c, "rsa_private"))

	//Decode base64 key
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(encodedKey)))
	base64.StdEncoding.Decode(base64Text, []byte(encodedKey))

	//Parse RSA key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(base64Text)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, helpers.ErrorWithCode("invalid_private", "Invalid private key!"))
		return
	}

	userInput := models.User{}
	if err := c.Bind(&userInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, helpers.ErrorWithCode("invalid_input", "Failed to bind the body data"))
		return
	}

	user, err := store.FindUser(c, params.M{"email": userInput.Email})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, helpers.ErrorWithCode("user_does_not_exist", "User does not exist"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, helpers.ErrorWithCode("incorrect_password", "Password is not correct"))
		return
	}

	token := jwt.New(jwt.GetSigningMethod(jwt.SigningMethodRS256.Alg()))
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["id"] = user.Id

	token.Claims = claims
	tokenString, err := token.SignedString(privateKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user.Sanitize()})
}
