package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"omikuji-app/pkg/api/ocontext"
)

func TestContextMiddleWare_ServeNext(t *testing.T) {
	type args struct {
		h http.Handler
	}
	tests := []struct {
		name string
		m    ContextMiddleWare
		args args
		want string
	}{
		{
			name: "リクエスト起因データの設定",
			m:    ContextMiddleWare{},
			args: args{
				h: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// Contextにセットされたアクセス時刻をレスポンスに書き込む
					fmt.Fprintf(w, "%s", ocontext.GetAccessTime(r.Context()).Format("2006-01-02 15:04:05"))
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.ServeNext(tt.args.h)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			got.ServeHTTP(w, r)

			rw := w.Result()
			defer rw.Body.Close()
			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// 時刻としてパースできればOKとする
			if _, err := time.Parse("2006-01-02 15:04:05", string(b)); err != nil {
				t.Errorf("parse error: %v", err)
			}
		})
	}
}

func TestResponseHeaderMiddleWare_ServeNext(t *testing.T) {
	type args struct {
		h http.Handler
	}
	tests := []struct {
		name string
		m    ResponseHeaderMiddleWare
		args args
		want string
	}{
		{
			name: "Content-Typeの設定",
			m:    ResponseHeaderMiddleWare{},
			args: args{
				h: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			want: "application/json; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.ServeNext(tt.args.h)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			got.ServeHTTP(w, r)
			if c := w.Header().Get("Content-Type"); c != tt.want {
				t.Errorf("Content-Type = %v, want %v", c, tt.want)
			}

		})
	}
}
