package omikuji

import (
	"context"
	"reflect"
	"testing"
	"time"

	"omikuji-app/pkg/api/ocontext"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want OmikujiService
	}{
		{
			name: "omikujiServiceの生成",
			want: &omikujiService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_omikujiService_Draw(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		s    *omikujiService
		args args
		want string
	}{
		{
			name: "正月以外のおみくじ",
			s:    &omikujiService{},
			args: args{
				ctx: ocontext.SetAccessTime(context.Background(), time.Date(2000, 1, 4, 3, 4, 5, 6, time.Local)),
			},
		},
		{
			name: "1/2のおみくじ",
			s:    &omikujiService{},
			args: args{
				ctx: ocontext.SetAccessTime(context.Background(), time.Date(2000, 1, 2, 3, 4, 5, 6, time.Local)),
			},
			want: "大吉",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Draw(tt.args.ctx)
			if tt.want != "" && !reflect.DeepEqual(got.Ruck, tt.want) {
				t.Errorf("omikujiService.Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}
