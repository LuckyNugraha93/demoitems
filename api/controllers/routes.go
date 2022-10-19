package controllers

import "github.com/LuckyNugraha93/demoitems/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Items routes
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateItem))).Methods("POST")
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetItems))).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetItem))).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateItem))).Methods("PUT")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteItem))).Methods("POST")

	//Transactions routes
	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateTransaction))).Methods("POST")
	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetTransactions))).Methods("GET")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetTransaction))).Methods("GET")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateTransaction))).Methods("PUT")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteTransaction))).Methods("POST")
}