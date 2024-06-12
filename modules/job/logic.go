package job

import (
	"dans_golang/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	description := r.URL.Query().Get("description")
	location := r.URL.Query().Get("location")
	fullTime := r.URL.Query().Get("full_time")
	page := r.URL.Query().Get("page")

	url := "https://dev6.dansmultipro.com/api/recruitment/positions.json"
	var queryParams []string

	if description != "" {
		queryParams = append(queryParams, "description="+description)
	}
	if location != "" {
		queryParams = append(queryParams, "location="+location)
	}
	if fullTime != "" {
		queryParams = append(queryParams, "full_time="+fullTime)
	}
	if page != "" {
		queryParams = append(queryParams, "page="+page)
	}

	if len(queryParams) > 0 {
		url += "?" + strings.Join(queryParams, "&")
	}

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Internal Server Error",
		})
		return
	}
	defer resp.Body.Close()

	var jobs []Job
	if err := json.NewDecoder(resp.Body).Decode(&jobs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Internal Server Error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helpers.Response{
		Message: "Success Get All Job",
		Data:    jobs,
	})
}

func GetJobDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	url := "https://dev6.dansmultipro.com/api/recruitment/positions/" + id

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Internal Server Error",
		})
		return
	}
	defer resp.Body.Close()

	var job Job
	if err := json.NewDecoder(resp.Body).Decode(&job); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Internal Server Error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helpers.Response{
		Message: "Success Get Detail Job",
		Data:    job,
	})
}
