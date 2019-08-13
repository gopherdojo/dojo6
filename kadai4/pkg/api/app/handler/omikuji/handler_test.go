package omikuji

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"omikuji-app/pkg/api/ocontext"

	service "omikuji-app/pkg/api/domain/service/omikuji"
	mockService "omikuji-app/pkg/api/domain/service/omikuji/mock"
)

func TestNew(t *testing.T) {
	type args struct {
		s service.OmikujiService
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		{
			name: "omikujiHandlerの生成",
			args: args{
				s: mockService.New(),
			},
			want: &omikujiHandler{omikujiService: mockService.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_omikujiHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		omikujiService service.OmikujiService
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
				omikujiService: mockService.New(),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil).WithContext(ocontext.SetAccessTime(context.Background(), time.Now())),
			},
			want:     "{\"id\":4,\"ruck\":\"吉\",\"message\":\"吉です！良い運勢ですね！\"}\n",
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &omikujiHandler{
				omikujiService: tt.fields.omikujiService,
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
