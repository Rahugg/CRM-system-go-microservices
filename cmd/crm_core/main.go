package main

import (
	"crm_system/config/crm_core"
	"crm_system/internal/crm_core/app"
)

// @title crm-core-service
// @version 1.0
// @description crm-core-service
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes https http
// @host localhost:8082
// @securityDefinitions.apikey	BearerAuth
// @type apiKey
// @name Authorization
// @in header
func main() {
	// Configuration
	cfg := crm_core.NewConfig()

	// Run
	app.Run(cfg)
}
