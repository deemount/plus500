package app

import (
	"github.com/gorilla/mux"

	"github.com/deemount/plus500/api/config"
)

// App ...
type App struct {
	// refer
	Routes config.Routes

	// pointer
	API     *config.API
	Plus500 *config.Plus500

	// legal
	Router *mux.Router
	V1     *mux.Router
}
