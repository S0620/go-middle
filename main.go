package main

import (
	"log"
	"fmt"
	"database/sql"
	"os"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"myapi/api"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn	   = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	
	r := api.NewRouter(db)//a


	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
 
}
