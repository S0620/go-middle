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
    row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	
	//ここまで
	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
	    select nice
		from articles
		where article_id = ?;
	`

    const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

 // (問4) 指定されたIDの記事のいいね数を+1するようにデータベースの中身を更新する処理
    tx, err := db.Begin()
	if err != nil {
		return err
	}

	var nicenum int
	row := tx.QueryRow(sqlGetNice, articleID)
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	if err := tx.Commit(); err != nil {
		return err
	}

 //ここまで
    return nil
}

//新規投稿をデータベースにinsertする関数
//->データベースに保存したコメント内容と、発生したエラーを返り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
	    insert into comments (article_id, message, created_at) values
		(?, ?, now())
	`

	 //(問5)構造体models.Commentを受け取って、それをデータベースに挿入する処理
	 result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	 if err != nil {
		return models.Comment{}, err
	 }
	 newID, err := result.LastInsertId()
	 if err != nil {
		return models.Comment{}, err
	 }
	 newComment := comment
	 newComment.CommentID = int(newID)
	 //ここまで
	return newComment, nil
}



