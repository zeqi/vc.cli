package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"vc.cli/routes"
	"vc.cli/services"
	"vc.cli/utils"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	instance := utils.GetInstance()
	config := instance.Config
	server := config.Server

	services.Init(config.MicroServices)

	router := gin.Default()
	router.Use(sessions.Sessions("vc.cli-session", routes.GetSesssionStore()))
	router.Use(static.Serve("/", static.LocalFile(server.Static, true)))
	routes.Handlers(server.Version, config.BasicAuth, router)
	router.Run(server.Host + ":" + strconv.Itoa(server.Port))
}
