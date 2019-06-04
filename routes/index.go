package routes

import (
	"github.com/gin-gonic/gin"
	"vc.cli/models"

	// "vc.cli/routes/gateway"
	"vc.cli/routes/gateway"
	"vc.cli/routes/sequence"
)

// Handlers for route management
func Handlers(version string, basicAuth models.BasicAuth, router *gin.Engine) {
	gateway.Register(router)

	// Login status validation
	group := router.Group(version, Permission())

	sequence.Register(group)

	// Basic Auth validation
	accounts := gin.Accounts{}
	accounts = basicAuth.Accounts
	basePath := basicAuth.Prefix + "/" + basicAuth.Version
	basicAuthGroup := router.Group(basePath, gin.BasicAuth(accounts))

	sequence.Register(basicAuthGroup)
}
