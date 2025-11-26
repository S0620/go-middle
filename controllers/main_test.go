package controllers_test

import (

	"myapi/controllers"
	"myapi/controllers/testdata"

	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テストに使うリソース（コントローラ構造体）を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}