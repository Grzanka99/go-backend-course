package utils

import (
	"encoding/json"
	"net/http"
)

func createResponse[T any | []any](msg T, rw http.ResponseWriter, statusCode int) int {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(msg)

	return statusCode
}

func Success[T any | []any](msg T, rw http.ResponseWriter) int {
	return createResponse(msg, rw, http.StatusOK)
}

func BadRequest[T any | []any](msg T, rw http.ResponseWriter) int {
	return createResponse(msg, rw, http.StatusBadRequest)
}

func InternalServerError[T any | []any](msg T, rw http.ResponseWriter) int {
	return createResponse(msg, rw, http.StatusInternalServerError)
}
