package repositories

import (
	"database/sql"

	"myapi/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
    if err != nil {
		return models.Article{}, err
	}
	id, _ :=result.LastInsertId()
	newArticle.ID = int(id)
	
    return newArticle, nil
}