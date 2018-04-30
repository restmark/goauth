package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/restmark/goauth/server"
)

func main() {
	api := &server.API{Router: gin.Default(), Config: viper.New()}

	// Configuration setup
	err := api.SetupViper()
	if err != nil {
		panic(err)
	}

	// Broker setup
	err = api.SetupProducer()
	if err != nil {
		panic(err)
	}

	// Router setup
	api.SetupRouter()
	api.Router.Run(api.Config.GetString("host_address"))
}
