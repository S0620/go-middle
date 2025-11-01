package services

import (
	"myapi/models"
	"myapi/repositories"

)
// PostCommentHandlerで使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func PostCommentService(comment models.Comment) (models.Comment, error) {
	// TODO:実装
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db,comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
