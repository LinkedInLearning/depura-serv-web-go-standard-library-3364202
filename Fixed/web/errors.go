package web

import "net/http"

func badRequestErrorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 Bad Request Error: " + err.Error()))
}

func internalServerErrorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error: " + err.Error()))
}
