package controllers

import (
	"net/http"

	"taktyl.com/m/src/api/responses"
)

// Home : /home endpoint
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
