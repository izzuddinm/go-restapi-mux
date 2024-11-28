package helper

import (
	"encoding/json"
	"net/http"
)

// BaseResponse is a structure for standardized API responses
type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"` // Optional field for error messages or extra details
}

// ResponseJson sends a standard JSON response with a given payload
func ResponseJson(w http.ResponseWriter, code int, payload BaseResponse) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ResponseError sends a standard error response
func ResponseError(w http.ResponseWriter, code int, message string) {
	payload := BaseResponse{
		Success: false,
		Message: message,
	}
	ResponseJson(w, code, payload)
}

// ResponseSuccess sends a successful response with data
func ResponseSuccess(w http.ResponseWriter, code int, data interface{}) {
	payload := BaseResponse{
		Success: true,
		Data:    data,
	}
	ResponseJson(w, code, payload)
}

// ResponseSuccessWithMessage sends a successful response with message and data
func ResponseSuccessWithMessage(w http.ResponseWriter, code int, message string) {
	payload := BaseResponse{
		Success: true,
		Message: message,
	}
	ResponseJson(w, code, payload)
}
