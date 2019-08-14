package mock

import (
	"fmt"

	"omikuji-app/pkg/api/app/presenter"
)

type mockPresenter struct {
}

func New() presenter.Presenter {
	return &mockPresenter{}
}

func (p *mockPresenter) Output(v interface{}) (string, error) {
	return fmt.Sprintf("output: %v", v), nil
}

type mockErrorPresenter struct {
}

func NewError() presenter.Presenter {
	return &mockErrorPresenter{}
}

func (p *mockErrorPresenter) Output(v interface{}) (string, error) {
	return "", &MockError{}
}

type MockError struct{}

func (e *MockError) Error() string {
	return "mock presenter error."
}
