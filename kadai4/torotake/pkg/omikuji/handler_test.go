package omikuji_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo6/kadai4/torotake/pkg/omikuji"
)

func TestServer_Handler(t *testing.T) {
	cases := []struct {
		name             string           // テスト名
		getTimeFunc      func() time.Time // Serverに渡す時刻取得関数
		expectedAllMatch bool             // 全施行のおみくじ結果が一致することを正とするかどうか
		expectedFortune  string           // 全思考のおみくじ結果が一致することを期待する場合、その結果
	}{
		{
			name:             "正月期間のときは全部大吉 開始境界 (1/1 00:00:00)",
			getTimeFunc:      func() time.Time { return time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local) },
			expectedAllMatch: true,
			expectedFortune:  "大吉",
		},
		{
			name:             "正月期間のときは全部大吉 終了境界 (1/3 23:59:59.999999999)",
			getTimeFunc:      func() time.Time { return time.Date(2019, time.January, 3, 23, 59, 59, 999999999, time.Local) },
			expectedAllMatch: true,
			expectedFortune:  "大吉",
		},
		{
			name:             "正月期間のときは全部大吉 開始境界直前 (12/31 23:59:59.999999999)",
			getTimeFunc:      func() time.Time { return time.Date(2018, time.December, 31, 23, 59, 59, 999999999, time.Local) },
			expectedAllMatch: false,
			expectedFortune:  "",
		},
		{
			name:             "正月期間のときは全部大吉 終了境界直後 (1/4 00:00:00)",
			getTimeFunc:      func() time.Time { return time.Date(2019, time.January, 4, 0, 0, 0, 0, time.Local) },
			expectedAllMatch: false,
			expectedFortune:  "",
		},
		{
			// TODO : 本当の正月に実行したらひっかかってしまう…
			name:             "正月期間以外の時にランダム",
			getTimeFunc:      nil,
			expectedAllMatch: false,
			expectedFortune:  "",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			runHandlerTest(t, c.getTimeFunc, c.expectedAllMatch, c.expectedFortune)
		})
	}
}

func runHandlerTest(t *testing.T, getTimeFunc func() time.Time, expectedAllMatch bool, expectedFortune string) {
	t.Helper()
	n := 100
	var first string
	var detectRandom bool
	for i := 0; i < n; i++ {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		s := omikuji.Server{
			GetTimeFunc: getTimeFunc,
		}
		s.Handler(w, r)
		rw := w.Result()
		defer rw.Body.Close()

		if rw.StatusCode != http.StatusOK {
			t.Errorf("unexpected status code")
		}

		b, err := ioutil.ReadAll(rw.Body)
		if err != nil {
			t.Errorf("unexpected error : reading response body failed")
		}

		var d = map[string]interface{}{}
		err = json.Unmarshal(b, &d)
		if err != nil {
			t.Errorf("unexpected error : unmarshaling json failed")
		}

		fortune, _ := d["fortune"].(string)
		if expectedAllMatch {
			// 全一致期待のときは期待値と違うのが返ってきた時点でエラー
			if fortune != expectedFortune {
				t.Fatalf("unexpected error : loop=%d, expect=%s, actual=%s", i, expectedFortune, fortune)
			}
		} else {
			// ランダム期待のときは全部結果が一緒であればエラー
			if i == 0 {
				// 初回の値を記憶
				first = fortune
			} else if i == n-1 {
				// 最後にチェック
				if !detectRandom {
					// 全部結果が一緒でランダムではなかった
					// 本当にランダムで試行が全部同じ値になったときは諦める
					t.Errorf("unexpected error : loop=%d, all result is same, not random. actual=%s", i, fortune)
				}
			} else {
				// 違うのが出てきたらランダムという事にする
				if fortune != first {
					detectRandom = true
				}
			}
		}
	}
}
