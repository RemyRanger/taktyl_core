package controllers

import (
	"taktyl.com/m/src/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Events routes
	s.Router.HandleFunc("/events", middlewares.SetMiddlewareJSON(s.CreateEvent)).Methods("POST")
	s.Router.HandleFunc("/events", middlewares.SetMiddlewareJSON(s.GetEvents)).Methods("GET")
	s.Router.HandleFunc("/events/{id}", middlewares.SetMiddlewareJSON(s.GetEvent)).Methods("GET")
	s.Router.HandleFunc("/events/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateEvent))).Methods("PUT")
	s.Router.HandleFunc("/events/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteEvent)).Methods("DELETE")
}
