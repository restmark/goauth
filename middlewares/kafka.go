package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/services"
)

func KafkaMiddleware(kafka services.KafkaInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(services.KafkaKey, kafka)
		c.Next()
	}
}
