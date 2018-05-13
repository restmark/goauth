package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/services"
	"github.com/spf13/viper"
	"github.com/globalsign/mgo"
)

type API struct {
	Router   *gin.Engine
	Config   *viper.Viper
	Kafka    services.KafkaInterface
	Database *mgo.Database
}
