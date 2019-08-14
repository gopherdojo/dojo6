package ocontext

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestSetAccessTime(t *testing.T) {
	type args struct {
		ctx  context.Context
		time time.Time
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "アクセス時刻の設定",
			args: args{
				ctx:  context.Background(),
				time: time.Date(2000, 1, 2, 3, 4, 5, 6, time.Local),
			},
			want: context.WithValue(context.Background(), accessTimeKey{}, time.Date(2000, 1, 2, 3, 4, 5, 6, time.Local)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetAccessTime(tt.args.ctx, tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAccessTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccessTime(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "アクセス時刻の取得",
			args: args{
				ctx: context.WithValue(context.Background(), accessTimeKey{}, time.Date(2000, 1, 2, 3, 4, 5, 6, time.Local)),
			},
			want: time.Date(2000, 1, 2, 3, 4, 5, 6, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAccessTime(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccessTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
