package main

import (
	"dans_golang/helpers"
	"dans_golang/modules/job"
	"dans_golang/modules/user"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize MySQL database
	helpers.InitDatabase()

	// Create router
	r := mux.NewRouter()

	// Route handlers
	user.Initiate(r)
	job.Initiate(r)

	// Start server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
