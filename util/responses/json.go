package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type returnData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type errData struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}

// JSON .
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(returnData{
		Status: statusCode,
		Data:   data,
	})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// Error .
func Error(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	errJson := json.NewEncoder(w).Encode(errData{
		Status: statusCode,
		Error:  err,
	})
	if errJson != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
