package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/restmark/goauth/config"
	"net/http"
)

type StatusController struct {
}

func NewStatusController() StatusController {
	return StatusController{}
}

func (s StatusController) GetApiStatus(c *gin.Context) {
	version := config.GetString(c, "version")
	c.JSON(http.StatusOK, gin.H{"status": "success", "version": version, "message": "You successfully reached the account API."})
}
