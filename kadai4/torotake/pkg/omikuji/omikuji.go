/*
Package omikuji はおみじくサーバーの機能を提供します。
サーバー実行環境のLocalの時刻で1/1〜1/3は必ず大吉が返ります。

おみじく抽選部分の実装
*/
package omikuji

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// Omikuji おみくじの結果を表す型
type omikuji struct {
	Fortune string `json:"fortune"`
	Message string `json:"message"`
}

// おみくじの定義リスト 0番目を大吉とする
var omikujiList = []omikuji{
	{"大吉", "今のあなたは運がいい！今なら何でも出来る…かも？"},
	{"吉", "結構ついてます。中吉より上だよ、知ってた？"},
	{"中吉", "何事もほどほど。運もほどほど。"},
	{"小吉", "小さい幸せを噛み締めましょう。"},
	{"末吉", "末広がり！後々良いことあるかもよ。"},
	{"凶", "しばらく大人しくしておいた方がいいかも……？"},
	{"大凶", "これ以上悪くなることはないよ。どんまい！"},
}

func init() {
	// 乱数シードの初期化
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func draw(t time.Time) omikuji {
	// 1/1〜1/3は大吉固定
	if isDaikichiDay(t) {
		return omikujiList[0]
	}

	// 通常はランダム選択
	index := rand.Intn(len(omikujiList))
	return omikujiList[index]
}

func isDaikichiDay(t time.Time) bool {
	// 1/1〜1/3は大吉固定
	if t.Month() == time.January && (t.Day() >= 1 && t.Day() <= 3) {
		return true
	}
	return false
}
