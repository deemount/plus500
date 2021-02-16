package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/deemount/plus500/api/config"
	"github.com/deemount/plus500/api/constants"
	"github.com/deemount/plus500/api/controllers"
)

// server is a global var
var server = controllers.Server{}

// Run is a func
// returns nil on success
func Run() error {

	// assign error
	var err error

	// Notice:
	// Put in comments when in production mode
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	//
	version, _ := strconv.Atoi(os.Getenv("API_PLUS500_VERSION"))
	limitClient, _ := strconv.Atoi(os.Getenv("API_SERVER_CLIENT_LIMIT"))

	server.App.API = &config.API{
		Version:     version,
		Host:        os.Getenv("API_SERVER_HOST"),
		Path:        os.Getenv("API_SERVER_PATH_PREFIX"),
		Port:        os.Getenv("API_SERVER_PORT"),
		Destination: os.Getenv("API_SERVER_PATH_SRC"),
		LimitClient: limitClient,
	}

	// Initialize routes
	server.App.Routes = config.Routes{
		config.Route{
			Name:        "Home",            // Name
			Method:      http.MethodGet,    // HTTP method
			Pattern:     constants.HOMEURI, // Route pattern
			HandlerFunc: server.Home,       // Handler function
		},
		config.Route{
			Name:        "HealthCheck",
			Method:      http.MethodGet,
			Pattern:     constants.HEALTHURI,
			HandlerFunc: server.HealthCheck,
		},
	}

	server.App.Plus500 = &config.Plus500{
		URL:       os.Getenv("API_PLUS500_URL"),
		UserAgent: os.Getenv("API_PLUS500_USERAGENT"),
	}

	// initialize orm idle and router
	if err = server.Initialize(); err != nil {
		return err
	}

	// run server
	if err = server.Run(); err != nil {
		return err
	}

	return nil

}
