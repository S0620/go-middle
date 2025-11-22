package apperrors

type MyAppError struct {
	ErrCode			//レスポンズとログに表示するエラーコード
	Message string  //レスポンズに表示するエラーメッセージ
	Err     error   `json:"-"` //エラーチェーンのための内部エラー
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
