package controllers

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	ErrorMessage string `json:"error_message"`
	Errors		 map[string]string `json:"errors,omitempty"`
}

// Send Response
func sendResponse(w http.ResponseWriter, content []byte, statusCode int, headers map[string]string) {
	for h, v := range headers {
		w.Header().Set(h, v)
	}

	w.WriteHeader(statusCode)
	w.Write(content)
}

// Send JSON response
func sendJSON(w http.ResponseWriter, data []byte, statusCode int) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	sendResponse(w, data, statusCode, headers)
}

// Marshal a struct and send as JSON
func sendStructAsJSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, _ := json.Marshal(payload)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	sendResponse(w, response, statusCode, headers)
}
