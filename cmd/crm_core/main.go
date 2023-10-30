package main

import (
	"crm_system/config/crm_core"
	"crm_system/internal/crm_core/app"
)

func main() {
	// Configuration
	cfg := crm_core.NewConfig()

	// Run
	app.Run(cfg)
}
