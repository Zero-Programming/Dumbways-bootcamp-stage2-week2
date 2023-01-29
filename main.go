package main

import (
	"fmt"
	"housy/database"
	"housy/pkg/mysql"
	"housy/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
