package helper

import (
	"encoding/json"
	"net/http"
)

//ResponseHandler generate response object
func ResponseHandler(w http.ResponseWriter, data interface{}, message string, httpstatus int) {

	response := make(map[string]interface{})
	sucess := false
	if httpstatus == 200 {
		sucess = true
	}
	response["status"] = sucess
	response["message"] = message
	response["data"] = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpstatus)

	json.NewEncoder(w).Encode(response)
}
