package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/config"
	"github.com/restmark/goauth/helpers"
	"github.com/restmark/goauth/models"
	"github.com/restmark/goauth/services"
	"github.com/restmark/goauth/store"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, helpers.ErrorWithCode("invalid_input", "Failed to bind the body data"))
		return
	}

	err = services.GetKafka(c).SendValue(user, config.GetString(c, "kafka_topic"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, helpers.ErrorWithCode("kafka_error", "Failed to send the data to Kafka "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (uc UserController) GetUser(c *gin.Context) {
	user, err := store.FindUserById(c, c.Param("id"))

	if err != nil {
		c.AbortWithError(http.StatusNotFound, helpers.ErrorWithCode("user_not_found", "The user does not exist"))
		return
	}

	c.JSON(http.StatusOK, user.Sanitize())
}
