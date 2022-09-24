package main

import (
	"fmt"
	"net/http"
	"prokdrn/database"
	"prokdrn/pkg/mysql"
	"prokdrn/routes"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// On http (API)
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
