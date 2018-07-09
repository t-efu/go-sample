package utils

import (
	"net/http"
)

// RenderInternalServerError - render internal server error
func RenderInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error: " + err.Error()))
	return
}

// RenderBadRequest - render bad request
func RenderBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	return
}
