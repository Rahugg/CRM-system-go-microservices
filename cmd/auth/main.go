package main

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/app"
)

func main() {
	// Configuration
	cfg := auth.NewConfig()

	// Run
	app.Run(cfg)
}
