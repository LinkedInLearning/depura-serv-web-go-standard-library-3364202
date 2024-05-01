package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/linkedinlearning/depura-go/webservice/web"
)

func mustInitDB() *sql.DB {
	user := os.Getenv("APP_USER")
	password := os.Getenv("APP_PASSWORD")
	dbname := os.Getenv("APP_DBNAME")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
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

	router.GET("/get-usr-nfo", userHandler.Info)
	router.Run(":8080")
}
