package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/linkedinlearning/depura-go/webservice/web"
)

func mustInitDB() *sql.DB {
	connStr := "user=user dbname=dbname password=passw4rd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	db := mustInitDB()

	var (
		host = flag.String("host", "", "host http address to listen on")
		port = flag.String("port", "8080", "port number for http listener")
	)
	flag.Parse()

	addr := net.JoinHostPort(*host, *port)

	if err := runHttp(addr, db); err != nil {
		log.Fatal(err)
	}
}

func runHttp(listenAddr string, db *sql.DB) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: web.NewRouter(db),
	}

	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)

	return s.ListenAndServe()
}
