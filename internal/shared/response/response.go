package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error ErrorDetail `json:"error""`
}

type ErrorDetail struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Detail  []FieldError `json:"detail"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// JSON envia resposta JSON com status code
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, statusCode int, code, message string) {
	response := ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
	JSON(w, statusCode, response)
}

func ValidationError(w http.ResponseWriter, details []FieldError) {
	response := ErrorResponse{
		Error: ErrorDetail{
			Code:    "VALIDATION_ERROR",
			Message: "Dados Inválidos",
			Detail:  details,
		},
	}
	JSON(w, http.StatusBadRequest, response)
}
