package helper

import (
	"encoding/json"
	"net/http"
)

//ResponseHandler generate response object
func ResponseHandler(w http.ResponseWriter, data interface{}, error map[string]interface{}, httpstatus int) {
	response := make(map[string]interface{})

	response["data"] = data
	response["error"] = error

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpstatus)

	json.NewEncoder(w).Encode(response)
}

//ApiErrorNew generate error object
func ApiErrorNew(httpStatus int, errorMessage string) map[string]interface{} {
	errorObject := make(map[string]interface{})
	errorObject["code"] = httpStatus
	errorObject["message"] = errorMessage

	return errorObject
}
