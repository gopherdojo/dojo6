package presenter

import (
	"reflect"
	"testing"

	entity "omikuji-app/pkg/api/domain/entity/omikuji"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Presenter
	}{
		{
			name: "presenterの生成",
			want: &presenter{},
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

func Test_presenter_Output(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		p       *presenter
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Jsonエンコード",
			p:    &presenter{},
			args: args{
				v: entity.OmikujiResult{
					ID:      1,
					Ruck:    "大吉",
					Message: "メッセージ1",
				},
			},
			want:    "{\"id\":1,\"ruck\":\"大吉\",\"message\":\"メッセージ1\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &presenter{}
			got, err := p.Output(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("presenter.Output() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("presenter.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
