package main

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/app"
)

// @title auth-service
// @version 1.0
// @description auth_service
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes https http
// @host localhost:8081
// @securityDefinitions.apikey	BearerAuth
// @type apiKey
// @name Authorization
// @in header
func main() {
	// Configuration
	cfg := auth.NewConfig()

	// Run
	app.Run(cfg)
}
