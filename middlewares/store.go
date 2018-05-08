package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/store"
	"github.com/restmark/goauth/store/mock"
	"github.com/restmark/goauth/store/mongodb"
	mgo "gopkg.in/mgo.v2"
)

func StoreMiddleware(db *mgo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.ToContext(c, mongodb.New(db))
		c.Next()
	}
}

func StoreMockMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		store.ToContext(c, mock.New())
		c.Next()
	}
}
