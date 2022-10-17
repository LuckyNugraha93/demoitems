package controllers

import "github.com/LuckyNugraha93/demoitems/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Items routes
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.CreateItem)).Methods("POST")
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.GetItems)).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(s.GetItem)).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateItem))).Methods("PUT")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(s.DeleteItem)).Methods("POST")

	//Transactions routes
	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(s.CreateTransaction)).Methods("POST")
	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(s.GetTransactions)).Methods("GET")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(s.GetTransaction)).Methods("GET")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateTransaction))).Methods("PUT")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(s.DeleteTransaction)).Methods("POST")
}