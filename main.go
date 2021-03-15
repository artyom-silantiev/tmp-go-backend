package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/artyom-silantiev/tmp-go-backend/lib/config"
	"github.com/artyom-silantiev/tmp-go-backend/lib/controllers"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println("Config: ", cfg)

	r := createRouter()
	r.Run(":" + strconv.Itoa(cfg.Port))
}

func createRouter() *gin.Engine {
	cfg := config.GetConfig()

	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	// public root
	if cfg.RootPublic != "" {
		fmt.Println("Public root: ", cfg.RootPublic)
		router.Use(static.Serve("/", static.LocalFile(cfg.RootPublic, false)))
	}

	// frontend admin spa root
	if cfg.RootFrontendAdminSpa != "" {
		fmt.Println("Frontend admin spa root: ", cfg.RootPublic)
		router.Use(static.Serve("/admin", static.LocalFile(cfg.RootFrontendAdminSpa, false)))
		router.Use(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/admin") && c.Request.Method == "GET" {
				c.File(cfg.RootFrontendAdminSpa + "index.html")
			}
		})
	}

	// frontend spa root
	if cfg.RootFrontendSpa != "" {
		fmt.Println("Frontend spa root: ", cfg.RootPublic)
		router.Use(static.Serve("/", static.LocalFile(cfg.RootFrontendSpa, false)))
		router.NoRoute(func(c *gin.Context) {
			c.File(cfg.RootFrontendSpa + "index.html")
		})
	}

	return router
}
