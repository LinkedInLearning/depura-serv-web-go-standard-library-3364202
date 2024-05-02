package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/linkedinlearning/depura-go/webservice/user"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) Info(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	userID, err := strconv.Atoi(uid)
	if err != nil {
		badRequestErrorHandler(w, err)
		return
	}

	ui, err := user.GetInfo(h.DB, userID)
	if err != nil {
		internalServerErrorHandler(w, err)
		return
	}

	// Set the status code to 201
	w.WriteHeader(http.StatusCreated)

	jsonBytes, err := json.Marshal(ui)
	if err != nil {
		log.Printf("Cannot encode User Info to JSON: %v", err)
		internalServerErrorHandler(w, err)
		return
	}
	w.Write(jsonBytes)
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/get-user-info" || r.Method == http.MethodGet {
		h.Info(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 Method not Allowed"))
}
