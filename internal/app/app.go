package app

import (
	"api_gateway/internal/config"
	"api_gateway/internal/transport/http/routes"
	"log"
	"net/http"
)

func Run(configPath string) {
	conf, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	router := routes.Init(conf)
	log.Fatal(http.ListenAndServe(":"+conf.Http.Port, router))
}
