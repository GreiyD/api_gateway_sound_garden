package main

import (
	"api_gateway/internal/app"
)

const configPath = "/srv/src/app/configs/main.yaml"

func main() {
	app.Run(configPath)
}
