package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "main関数の実行"},
	}
	for _, tt := range tests {
		// 他のテストでパターンは網羅しているので正常終了のみ確認
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
