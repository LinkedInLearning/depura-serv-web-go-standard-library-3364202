package web

import (
	"database/sql"
	"net/http"
)

func NewRouter(sb *sql.DB) http.Handler {
	mux := http.NewServeMux()

	userHandlers := &UserHandler{
		DB: sb,
	}

	mux.Handle("/get-usr-nfo", userHandlers)

	return mux
}
