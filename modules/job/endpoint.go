package job

import (
	"dans_golang/middlewares"
	"github.com/gorilla/mux"
)

func Initiate(r *mux.Router) {
	r.Handle("/jobs", middlewares.JwtAuth(GetJobs)).Methods("GET")
	r.Handle("/jobs/{id}", middlewares.JwtAuth(GetJobDetail)).Methods("GET")
}
