package services

import "errors"

//新しいエラーを作る
var ErrNoData = errors.New("get 0 record from db.Query")