package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success code, Usually 200
	Code int
	//Account balance
	Balance int64
}

// Error Response
// Error Response
type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
