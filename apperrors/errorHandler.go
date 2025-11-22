package apperrors

import (
	"net/http"
	"errors"
	"encoding/json"
	
	
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	//エラーの種類を判別して、適切なhttpレスポンスを返す
	var appErr *MyAppError

	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}
	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
	

}
