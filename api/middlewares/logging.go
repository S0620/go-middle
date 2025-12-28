package middlewares

import (
	"context"
	"log"
	"net/http"
)

//自作ResponseWriterを作る
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

//resLoggingWriter構造体のコンストラクタ
//->内部フィールドに入れるResponseWriterを受け取ってresLoggingWriter構造体を作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

//WriteHeaderメソッドを作る
func (rsw *resLoggingWriter) WriteHeader(code int) {
	//resLoggingWriter構造体のcodeフィールドに、使うレスポンスコードを保存する
	rsw.code = code
	//HTTPレスポンスに使うレスポンスコードを指定
	//(=WriteHeaderメソッド本来の機能を呼び出し)
	rsw.ResponseWriter.WriteHeader(code)
}

//ミドルウェアの中身
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := NewTraceID()
		
		//リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := req.Context()
		ctx = context.WithValue(ctx, traceIDKey{}, traceID)
		req = req.WithContext(ctx)
		//自作のResponseWriterを作って
		rlw := NewResLoggingWriter(w)

		//それをハンドラに渡す
		next.ServeHTTP(rlw, req)

		//自作ResponseWriterからロギングしたいデータを出す
		log.Printf("[%d]res: %d", traceID, rlw.code)
	}) 
}