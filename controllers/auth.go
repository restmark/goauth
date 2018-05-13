package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/config"
	"github.com/restmark/goauth/helpers"
	"github.com/restmark/goauth/helpers/params"
	"github.com/restmark/goauth/models"
	"github.com/restmark/goauth/store"
	"github.com/square/go-jose"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController struct {
}

func NewAuthController() AuthController {
	return AuthController{}
}

func (ac *AuthController) Authentication(c *gin.Context) {
	userInput := models.User{}
	if err := c.Bind(&userInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, helpers.ErrorWithCode("invalid_input", "Failed to bind the body data", err))
		return
	}

	user, err := store.FindUser(c, params.M{"email": userInput.Email})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, helpers.ErrorWithCode("user_does_not_exist", "User does not exist", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, helpers.ErrorWithCode("incorrect_password", "Password is not correct", err))
		return
	}

	//Read base64 private key
	encodedKey := []byte(config.GetString(c, "rsa_private"))

	refreshToken, err := helpers.GenerateRefreshToken(encodedKey, user.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, helpers.ErrorWithCode("token_generation_failed", "Could not generate the refresh token", err))
		return
	}

	accessToken, err := helpers.GenerateAccessToken(encodedKey, user.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, helpers.ErrorWithCode("token_generation_failed", "Could not generate the access token", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"refresh": refreshToken, "token": accessToken, "user": user.Sanitize()})
}

// refreshes the jwt refresh token
func (ac *AuthController) RefreshToken(c *gin.Context) {
	type RefreshTokenParams struct {
		RefreshToken string `json:"refresh_token"`
	}

	//Read base64 private key
	encodedKey := []byte(config.GetString(c, "rsa_private"))

	var refreshTokenParams RefreshTokenParams
	if err := c.Bind(&refreshTokenParams); err != nil {
		c.AbortWithError(http.StatusBadRequest, helpers.ErrorWithCode("invalid_input", "Failed to bind the body data", err))
		return
	}

	claims, err := helpers.ValidateJwtToken(refreshTokenParams.RefreshToken, encodedKey, "refresh")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, helpers.ErrorWithCode("invalid_token", "the given token is invalid", err))
		return
	}

	accessToken, err := helpers.GenerateAccessToken(encodedKey, claims["sub"].(string))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, helpers.ErrorWithCode("token_generation_failed", "Could not generate the access token", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": accessToken})
}

// returns the public key as a JWK
func (ac *AuthController) GetJWKS(c *gin.Context) {
	encodedPrivateKey := config.FromContext(c).GetString("rsa_private")

	privateKey, err := helpers.GetRSAPrivateKey([]byte(encodedPrivateKey))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, helpers.ErrorWithCode("invalid_private_key", "Could not get the private key", err))
		return
	}

	jwkToken := jose.JsonWebKey{Key: &privateKey.PublicKey, KeyID: "LOGIN", Algorithm: "RS256", Use: "sig"}
	json, err := jwkToken.MarshalJSON()
	if err != nil {
		c.Error(helpers.NewErrorWithTrace(http.StatusInternalServerError, "json_generation_failed", "Could not generate the public key json", err))
		c.Abort()
		return
	}

	c.Data(http.StatusOK, http.DetectContentType(json), json)
}
