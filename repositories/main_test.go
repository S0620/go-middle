package repositories_test

import (
	"testing"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//テスト全体で共有するsql.DB型
var testDB *sql.DB

//全テスト共通の前処理
func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

//全テスト共通の後処理を書く
func teardown() {
	testDB.Close()
}

//
func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}