package main

import (
	"fmt"
	"strconv"

	"github.com/artyom-silantiev/tmp-go-backend/lib/config"
	"github.com/artyom-silantiev/tmp-go-backend/lib/routes"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println("Config: ", cfg)

	r := routes.CreateRouter(cfg)
	r.Run(":" + strconv.Itoa(cfg.PORT))
}
