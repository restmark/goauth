package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/server"
	"github.com/spf13/viper"
)

func main() {
	api := &server.API{Router: gin.Default(), Config: viper.New()}

	// Configuration setup
	err := api.SetupViper()
	if err != nil {
		panic(err)
	}

	// Database setup
	session, err := api.SetupDatabase()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Broker setup
	err = api.SetupProducer()
	if err != nil {
		panic(err)
	}

	api.SetupTopicRouter()

	// Consumer setup
	go api.SetupConsumer()

	// Router setup
	api.SetupRouter()
	api.Router.Run(api.Config.GetString("host_address"))
}
