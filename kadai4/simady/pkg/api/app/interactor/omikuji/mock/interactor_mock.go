package mock

import (
	"context"

	interactor "omikuji-app/pkg/api/app/interactor/omikuji"
)

type mockInteractor struct {
}

func New() interactor.OmikujiInteractor {
	return &mockInteractor{}
}

func (i *mockInteractor) Draw(ctx context.Context) (string, error) {
	return "{\"id\":4,\"ruck\":\"吉\",\"message\":\"吉です！良い運勢ですね！\"}\n", nil
}

type mockErrorInteractor struct {
}

func NewError() interactor.OmikujiInteractor {
	return &mockErrorInteractor{}
}

func (i *mockErrorInteractor) Draw(ctx context.Context) (string, error) {
	return "", &MockError{}
}

type MockError struct{}

func (e *MockError) Error() string {
	return "mock interactor error."
}
