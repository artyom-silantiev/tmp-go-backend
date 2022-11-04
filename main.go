package main

import (
	"strconv"

	"github.com/artyom-silantiev/tmp-go-backend/lib/config"
	"github.com/artyom-silantiev/tmp-go-backend/lib/routes"
)

func main() {
	cfg := config.GetConfig()
	config.PrintConfig(cfg)
	r := routes.CreateRouter(cfg)
	r.Run(":" + strconv.Itoa(cfg.Port))
}
