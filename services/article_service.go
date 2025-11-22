package services

import (
	"myapi/models"
	"myapi/repositories"
	"myapi/apperrors"
	"errors"
	"database/sql"
	
)

//ArticleDetailHandlerで使うことを想定したサービス
//指定IDの記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	
	//1.repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFiled.Wrap(err, "fail to get data")
		return models.Article{}, err
	}
	commentlist, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFiled.Wrap(err, "fail to get data")
		return models.Article{}, err
	}
	//3.2で得たコメント一覧を、1で得たArticle構造体に紐付ける
	article.CommentList = append(article.CommentList, commentlist...)

	return article, nil
}


//PostArticleHandlerで使うことを想定したサービス
//指定した記事をデータベースに挿入し、挿入後の記事情報を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

//ArticleListHandlerで使うことを想定したサービス
//指定pageの記事一覧を返却

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFiled.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err = apperrors.GetDataFiled.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

//PostNiceHandlerで使うことを想定したサービス
//指定IDの記事のいいね数を+1して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
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

