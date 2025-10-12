package repositories

import (
	"database/sql"

	"myapi/models"
)

const (
	articleNumPerPage = 5	
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`
    //ここから問１
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
    if err != nil {
		return models.Article{}, err
	}
	id, _ :=result.LastInsertId()
	newArticle.ID = int(id)
	//ここまで
	
    return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	    select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`
	//ここから問２
    rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	//ここまで
	return articleArray, nil
}

func SelectArticleDetail(db*sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
	    select *
		from articles
		where article_id = ?;
	`
	//ここから問３
   
	
	//ここまで
	return article, nil
}
