package gateway

import (
	"github.com/gin-gonic/gin"
)

// Handlers for route management
func Register(router *gin.Engine) {
	controller := NewController()
	router.GET("/", controller.App)
	router.GET("/profile", controller.Profile)
}
