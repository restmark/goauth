package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/controllers"
	"github.com/restmark/goauth/middlewares"
)

func (a *API) SetupRouter() {
	router := a.Router

	/* middlewares */
	router.Use(gin.Recovery())
	router.Use(middlewares.Logger())
	router.Use(middlewares.ErrorMiddleware())
	router.Use(middlewares.ConfigMiddleware(a.Config))
	router.Use(middlewares.KafkaMiddleware(a.Kafka))
	router.Use(middlewares.StoreMiddleware(a.Database))

	v1 := router.Group("/v1")
	{
		statusController := controllers.NewStatusController()
		v1.GET("/", statusController.GetApiStatus)

		users := v1.Group("/users")
		{
			userController := controllers.NewUserController()
			users.POST("/", userController.CreateUser)
		}

		authentication := v1.Group("/auth")
		{
			authController := controllers.NewAuthController()
			authentication.POST("/", authController.Authentication)
			authentication.GET("/cert", authController.GetJWKS)
			authentication.POST("/refresh", authController.RefreshToken)
		}
	}
}
