package ocontext

import (
	"context"
	"time"
)

type accessTimeKey struct{}

// SetAccessTime contextにアクセス日時を設定する.
func SetAccessTime(ctx context.Context, time time.Time) context.Context {
	return context.WithValue(ctx, accessTimeKey{}, time)
}

// GetAccessTime contextからアクセス日時を取得する.
func GetAccessTime(ctx context.Context) time.Time {
	return ctx.Value(accessTimeKey{}).(time.Time)
}
