package services

import (
	"myapi/models"
	"myapi/repositories"
	
)

//ArticleDetailHandlerで使うことを想定したサービス
//指定IDの記事情報を返却
func GetArticleService(articleID int) (models.Article, error) {
	
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	//1.repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	commentlist, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	//3.2で得たコメント一覧を、1で得たArticle構造体に紐付ける
	article.CommentList = append(article.CommentList, commentlist...)

	return article, nil
}