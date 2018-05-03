package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/services"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

type API struct {
	Router      *gin.Engine
	Config      *viper.Viper
	Kafka       services.KafkaInterface
	Database    *mgo.Database
	TopicRouter *TopicRouter
}
