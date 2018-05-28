package app

import (
	"encoding/json"
	"net/http"
)

// FileSearch handles the info and list HTTP requests
func FileSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect method for info request"))
		return
	}
	var request SearchRequest
	err := request.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect list request"))
		return
	}
	files := Retrieve(request)
	result, _ := json.Marshal(files)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
