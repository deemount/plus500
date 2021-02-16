package controllers

import (
	"net/http"

	"github.com/deemount/plus500/api/responses"
)

// Home Controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Plus500 API v1")
}
