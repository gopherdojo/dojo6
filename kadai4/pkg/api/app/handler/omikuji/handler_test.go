package omikuji

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	interactor "omikuji-app/pkg/api/app/interactor/omikuji"
	mockInteractor "omikuji-app/pkg/api/app/interactor/omikuji/mock"
	"omikuji-app/pkg/api/ocontext"
)

func TestNew(t *testing.T) {
	type args struct {
		i interactor.OmikujiInteractor
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		{
			name: "omikujiHandlerの生成",
			args: args{
				i: mockInteractor.New(),
			},
			want: &omikujiHandler{omikujiInteractor: mockInteractor.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_omikujiHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		omikujiInteractor interactor.OmikujiInteractor
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     string
		wantCode int
	}{
		{
			name: "リクエスト成功",
			fields: fields{
				omikujiInteractor: mockInteractor.New(),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil).WithContext(ocontext.SetAccessTime(context.Background(), time.Now())),
			},
			want:     "{\"id\":4,\"ruck\":\"吉\",\"message\":\"吉です！良い運勢ですね！\"}\n",
			wantCode: http.StatusOK,
		},
		{
			name: "リクエスト失敗",
			fields: fields{
				omikujiInteractor: mockInteractor.NewError(),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil).WithContext(ocontext.SetAccessTime(context.Background(), time.Now())),
			},
			want:     "抽選に失敗しました.\n",
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &omikujiHandler{
				omikujiInteractor: tt.fields.omikujiInteractor,
			}
			h.ServeHTTP(tt.args.w, tt.args.r)
			rw := tt.args.w.(*httptest.ResponseRecorder).Result()
			defer rw.Body.Close()
			if rw.StatusCode != tt.wantCode {
				t.Errorf("unexpected status code: %v, want %v", rw.StatusCode, tt.wantCode)
			}
			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if s := string(b); s != tt.want {
				t.Errorf("unexpected response: %v, want %v", s, tt.want)
			}
		})
	}
}
