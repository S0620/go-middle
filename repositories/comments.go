package repositories

import (
	"database/sql"
	"myapi/models"
)


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

//指定IDの記事についたコメント一覧を取得する関数
//->取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
	    select *
	    from comments
	    where article_id = ?
    `
	//(問6)指定IDの記事についたコメント一覧をデータベースから取得して、
    //それを`models.Comment`構造体のスライス`[]models.Comment`に詰めて返す処理
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
    for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
			return nil, err
		}
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}
	

	//ここまで
	return commentArray, nil
}
