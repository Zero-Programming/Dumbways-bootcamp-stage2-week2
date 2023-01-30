package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	// r.HandleFunc("/houses", h.FindHouses).Methods("GET")
	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	// r.HandleFunc("/house", h.CreateHouse).Methods("POST")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
}
