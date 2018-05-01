package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/config"
	"github.com/spf13/viper"
)

func ConfigMiddleware(viper *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.ToContext(c, config.New(viper))
		c.Next()
	}
}
