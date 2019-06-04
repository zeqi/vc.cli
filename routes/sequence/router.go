package sequence

import (
	"github.com/gin-gonic/gin"
)

// Handlers for route management
func Register(router *gin.RouterGroup) *gin.RouterGroup {

	controller := NewController()

	group := router.Group("sequences")
	group.POST("/", controller.Create)
	group.GET("/", controller.Find)
	group.GET("/one", controller.FindOne)
	group.GET("/id/:id", controller.FindById)
	group.PUT("/name/:name", controller.IncByName)
	return group
}
