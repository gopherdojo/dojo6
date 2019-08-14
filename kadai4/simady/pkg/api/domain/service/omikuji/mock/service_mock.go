package mock

import (
	"context"

	entity "omikuji-app/pkg/api/domain/entity/omikuji"
	"omikuji-app/pkg/api/domain/service/omikuji"
)

// mockService OmikujiServiceのモック.
type mockService struct {
}

// New モックを生成する.
func New() omikuji.OmikujiService {
	return &mockService{}
}

// Draw おみくじを引く.
func (s *mockService) Draw(ctx context.Context) entity.OmikujiResult {
	return entity.OmikujiResult{
		ID:      4,
		Ruck:    "吉",
		Message: "吉です！良い運勢ですね！",
	}
}
