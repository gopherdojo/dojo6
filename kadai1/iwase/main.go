package main

import (
	"./Converter"
)

func main() {
	// 変換用インスタンスを作成
	c := Converter.NewConverter()
	// 変換処理を実行
	c.Convert()
}
