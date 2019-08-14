package interactor

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"omikuji-app/pkg/api/app/presenter"
	mockPresenter "omikuji-app/pkg/api/app/presenter/mock"
	entity "omikuji-app/pkg/api/domain/entity/omikuji"
	service "omikuji-app/pkg/api/domain/service/omikuji"
	mockService "omikuji-app/pkg/api/domain/service/omikuji/mock"
)

func TestNew(t *testing.T) {
	type args struct {
		p presenter.Presenter
		s service.OmikujiService
	}
	tests := []struct {
		name string
		args args
		want OmikujiInteractor
	}{
		{
			name: "omikujiInteractorの生成",
			args: args{
				p: mockPresenter.New(),
				s: mockService.New(),
			},
			want: &omikujiInteractor{presenter: mockPresenter.New(), omikujiService: mockService.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.p, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_omikujiInteractor_Draw(t *testing.T) {
	type fields struct {
		presenter      presenter.Presenter
		omikujiService service.OmikujiService
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常終了",
			fields: fields{
				presenter:      mockPresenter.New(),
				omikujiService: mockService.New(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: fmt.Sprintf("output: %v", entity.OmikujiResult{
				ID:      4,
				Ruck:    "吉",
				Message: "吉です！良い運勢ですね！",
			}),
			wantErr: false,
		},
		{
			name: "エラー発生",
			fields: fields{
				presenter:      mockPresenter.NewError(),
				omikujiService: mockService.New(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &omikujiInteractor{
				presenter:      tt.fields.presenter,
				omikujiService: tt.fields.omikujiService,
			}
			got, err := i.Draw(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("omikujiInteractor.Draw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("omikujiInteractor.Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}
