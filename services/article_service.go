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


//PostArticleHandlerで使うことを想定したサービス
//指定した記事をデータベースに挿入し、挿入後の記事情報を返却
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newarticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newarticle, nil
}

//ArticleListHandlerで使うことを想定したサービス
//指定pageの記事一覧を返却

func GetArticleListService(page int) ([]models.Article, error) {
    db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return []models.Article{}, err
	}

	return articleList, nil
}

//PostNiceHandlerで使うことを想定したサービス
//指定IDの記事のいいね数を+1して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	},nil
}

