package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/restmark/goauth/middlewares"
	"github.com/restmark/goauth/models"
	"github.com/restmark/goauth/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_CreateUser(t *testing.T) {
	userController := NewUserController()

	/* SETUP MOCK USER */
	testUser := models.User{
		Id:       "123",
		Firstname: "maxence",
		Lastname: "henneron",
		Email:    "maxence@restmark.co",
	}

	userJson, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf("Couldn't create json: %v\n", err)
	}

	/* SETUP CONFIG MOCK */
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	var configYaml = []byte(`
		kafka_topic: foo
	`)

	viper.ReadConfig(bytes.NewBuffer(configYaml))

	/* SETUP GIN */
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, and register routes
	r := gin.Default()
	r.Use(middlewares.ConfigMiddleware(viper.GetViper()))
	r.Use(middlewares.KafkaMiddleware(&services.KafkaMock{}))
	r.POST("/", userController.CreateUser)

	/* CREATE POST REQUEST */
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(userJson))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, w.Code, http.StatusOK)
}
