package controllers_test

import (
	"database/sql"
	"fmt"
	"log"
	"myapi/controllers"
	"myapi/services"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テストに使うリソース（コントローラ構造体）を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("DB setup fail")
		os.Exit(1)
	}

	ser := services.NewMyAppService(db)
	aCon = controllers.NewArticleController(ser)

	m.Run()
}