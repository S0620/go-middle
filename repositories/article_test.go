package repositories_test

import (
	"testing"
	"myapi/repositories"
	"database/sql"
	"myapi/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expected := models.Article {
	    ID:       1,
        Title:    "firstPost",
        Contents: "This is my first blog",
        UserName: "saki",
        NiceNum : 3,
    }

	//テスト対象となる関数を実行

	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		//関数の実行そのものに失敗した場合、テストも失敗させる
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}

	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}

	if got.Contents != expected.Contents {
		t.Errorf("Content: get %s but want %s\n", got.Contents, expected.Contents)
	}

	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}

	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}

	//t.Fatalもt.Errorも実行されずに関数が終わった場合にはテスト成功
}
