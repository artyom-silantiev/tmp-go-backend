package routes

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/artyom-silantiev/tmp-go-backend/lib/config"
	"github.com/artyom-silantiev/tmp-go-backend/lib/controllers"
)

func CreateRouter(cfg config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	// public root
	if cfg.DirHttpPublic != "" {
		fmt.Println("Public root: ", cfg.DirHttpPublic)
		router.Use(static.Serve("/", static.LocalFile(cfg.DirHttpPublic, true)))
	}

	// admin app root
	if cfg.DirHttpAdmin != "" {
		fmt.Println("Admin app root: ", cfg.DirHttpAdmin)
		router.Use(static.Serve("/admin", static.LocalFile(cfg.DirHttpAdmin, true)))
	}

	return router
}
