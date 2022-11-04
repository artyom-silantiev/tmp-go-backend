package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port          int    `envconfig:"APP_PORT" default:"3000"`
	DirHttpPublic string `envconfig:"APP_DIR_HTTP_PUBLIC"`
	DirHttpAdmin  string `envconfig:"APP_DIR_HTTP_ADMIN"`
}

func GetConfig() Config {
	var s Config

	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return s
}

func PrintConfig(s Config) {
	fmt.Printf("Config:\n")
	fmt.Printf("Port: %v\n", s.Port)
	fmt.Printf("DirHttpPublic: %s\n", s.DirHttpPublic)
	fmt.Printf("DirHttpAdmin: %s\n", s.DirHttpAdmin)
	fmt.Printf("\n")
}
