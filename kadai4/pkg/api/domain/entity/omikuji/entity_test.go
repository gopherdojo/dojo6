package omikuji

import (
	"reflect"
	"testing"
)

func TestOmikujiResults_FindRandom(t *testing.T) {
	tests := []struct {
		name string
		rs   OmikujiResults
	}{
		{
			name: "ランダム選択",
			rs: OmikujiResults{
				{
					ID:      1,
					Ruck:    "大吉",
					Message: "メッセージ1",
				},
				{
					ID:      2,
					Ruck:    "吉",
					Message: "メッセージ2",
				},
				{
					ID:      3,
					Ruck:    "中吉",
					Message: "メッセージ3",
				},
				{
					ID:      4,
					Ruck:    "小吉",
					Message: "メッセージ4",
				},
				{
					ID:      5,
					Ruck:    "凶",
					Message: "メッセージ5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.rs.FindRandom()
			// ランダム値で正否を判断できないのでログだけ出力
			t.Log(res)
		})
	}
}

func TestOmikujiResults_ExtractByRuck(t *testing.T) {
	type args struct {
		ruck string
	}
	tests := []struct {
		name string
		rs   OmikujiResults
		args args
		want OmikujiResults
	}{
		{
			name: "小吉を抽出",
			rs: OmikujiResults{
				{
					ID:      1,
					Ruck:    "大吉",
					Message: "メッセージ1",
				},
				{
					ID:      2,
					Ruck:    "吉",
					Message: "メッセージ2",
				},
				{
					ID:      3,
					Ruck:    "小吉",
					Message: "メッセージ3",
				},
				{
					ID:      4,
					Ruck:    "小吉",
					Message: "メッセージ4",
				},
				{
					ID:      5,
					Ruck:    "凶",
					Message: "メッセージ5",
				},
			},
			args: args{
				ruck: "小吉",
			},
			want: OmikujiResults{
				{
					ID:      3,
					Ruck:    "小吉",
					Message: "メッセージ3",
				},
				{
					ID:      4,
					Ruck:    "小吉",
					Message: "メッセージ4",
				},
			},
		},
		{
			name: "一致するものがない",
			rs: OmikujiResults{
				{
					ID:      1,
					Ruck:    "大吉",
					Message: "メッセージ1",
				},
				{
					ID:      2,
					Ruck:    "吉",
					Message: "メッセージ2",
				},
			},
			args: args{
				ruck: "凶",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rs.ExtractByRuck(tt.args.ruck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OmikujiResults.ExtractByRuck() = %v, want %v", got, tt.want)
			}
		})
	}
}
