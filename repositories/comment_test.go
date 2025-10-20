package repositories_test

import (
	"testing"
	"myapi/repositories"
	"myapi/models"
	_ "github.com/go-sql-driver/mysql"
)

//SelectCommentList関数のテスト
func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got Id %d\n", articleID, comment.ArticleID)
		}
	}
}
//InsertComment関数のテスト
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID:   2,
		Message:     "Insert test",
	}

	expectedCommentNum := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new comment id is %d but got %d\n",expectedCommentNum, newComment.CommentID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where comment_id = ?;
		`
		const sqlReset = `
			alter table comments auto_increment = 3;
		`

		_, err := testDB.Exec(sqlStr, newComment.CommentID)
		if err != nil {
			t.Error(err)
		}
		_, err = testDB.Exec(sqlReset)
		if err != nil {
			t.Error(err)
		}
	})
}