package main

import (
	"api_gateway/internal/app"
)

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
