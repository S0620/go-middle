package services

import (
	"database/sql"
	"errors"
	"myapi/apperrors"
	"myapi/models"
	"myapi/repositories"

)

//ArticleDetailHandlerで使うことを想定したサービス
//指定IDの記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)
	//1.repositories層の関数SelectArticleDetailで記事の詳細を取得
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList *[]models.Comment //1つの記事に対してコメントは複数個つく可能性があるため*[]を加えている
		err error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(s.db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, s.db, articleID)

	//ゴールーチンをselect文で処理。forループを使用して2回実行している
	for i:= 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	
	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			articleGetErr = apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, articleGetErr
		}
		articleGetErr = apperrors.GetDataFiled.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, articleGetErr
	}
	//2.コメント一覧を取得
	if commentGetErr != nil {
		err := apperrors.GetDataFiled.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}
	//3.2で得たコメント一覧を、1で得たArticle構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)

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

