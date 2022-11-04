package routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/artyom-silantiev/tmp-go-backend/lib/config"
	"github.com/artyom-silantiev/tmp-go-backend/lib/controllers"
)

func CreateRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	// public root
	if cfg.ROOT_PUBLIC != "" {
		fmt.Println("Public root: ", cfg.ROOT_PUBLIC)
		router.Use(static.Serve("/", static.LocalFile(cfg.ROOT_PUBLIC, false)))
	}

	// frontend admin spa root
	if cfg.ROOT_FRONTEND_ADMIN_SPA != "" {
		fmt.Println("Frontend admin spa root: ", cfg.ROOT_PUBLIC)
		router.Use(static.Serve("/admin", static.LocalFile(cfg.ROOT_FRONTEND_ADMIN_SPA, false)))
		router.Use(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/admin") && c.Request.Method == "GET" {
				c.File(cfg.ROOT_FRONTEND_ADMIN_SPA + "index.html")
			}
		})
	}

	// frontend spa root
	if cfg.ROOT_FRONTEND_SPA != "" {
		fmt.Println("Frontend spa root: ", cfg.ROOT_PUBLIC)
		router.Use(static.Serve("/", static.LocalFile(cfg.ROOT_FRONTEND_SPA, false)))
		router.NoRoute(func(c *gin.Context) {
			c.File(cfg.ROOT_FRONTEND_SPA + "index.html")
		})
	}

	return router
}
