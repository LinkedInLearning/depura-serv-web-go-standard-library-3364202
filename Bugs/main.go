package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	userHandler := web.UserHandler{
		DB: db,
	}

	router.POST("/get-usr-nfo", userHandler.Info)
	router.Run(":8080")
}
