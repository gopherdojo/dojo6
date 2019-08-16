/*
Package omikuji はおみじくサーバーの機能を提供します。
サーバー実行環境のLocalの時刻で1/1〜1/3は必ず大吉が返ります。

HTTPハンドラ部分の実装
*/
package omikuji

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Server おみくじのHTTPサーバー用の構造体
type Server struct {
	// おみくじ抽選結果の元になる時刻を取得する関数。nilの場合は現在時刻が使われる
	GetTimeFunc func() time.Time
}

// Handler おみくじAPIのhttp handler
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	var t time.Time
	if s.GetTimeFunc != nil {
		// 時刻取得関数が提供されていればそちらを使う
		t = s.GetTimeFunc()
	} else {
		// デフォルトは現在時刻
		t = time.Now()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// おみくじ抽選
	lot := draw(t)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(lot); err != nil {
		// jsonエンコードに失敗 Internal Server Errorとして返す
		http.Error(w, "Internal Server Error", 500)
		return
	}

	fmt.Fprint(w, buf.String())
}
