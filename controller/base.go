package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func WriteSuccess(w http.ResponseWriter, statusCode int, response interface{}) (int, error) {
	result, _ := json.Marshal(Response{
		Code: statusCode,
		Data: response,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(result)

	return statusCode, nil
}

func WriteError(w http.ResponseWriter, statusCode int, err error) (int, error) {
	result, _ := json.Marshal(Response{
		Code:  statusCode,
		Error: err.Error(),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(result)

	return statusCode, err
}
