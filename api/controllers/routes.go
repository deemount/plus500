package controllers

import (
	"net/http"

	"github.com/deemount/plus500/api/constants"
)

// initializeRoutes is a method
func (server *Server) initializeRoutes() error {

	var err error

	// uri
	home := constants.HOMEURI

	//**************** Home Route

	err = server.App.V1.HandleFunc(home, server.Home).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	//**************** Market Routes

	// single request

	return nil

}
