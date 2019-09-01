# 課題3-2

## 分割ダウンローダを作ろう

- Rangeアクセスを用いる
- いくつかのゴルーチンでダウンロードしてマージする
- エラー処理を工夫する
  - golang.org/x/sync/errgroupパッケージなどを使ってみる
- キャンセルが発生した場合の実装を行う

## 使い方

```
go get github.com/gopherdojo/dojo6/kadai3-2/en-ken/dl-mgr
div -n [goroutineの並列数(デフォルト:5)] -o [保存ファイル名(デフォルト:リモート名)] URL
```

## やったこと

1. HEADリクエストしてAccept-Rangesヘッダの有無確認
1. Accept-Rangesがあった場合、
    1. 任意の分割数に応じて、goroutineごとの割当範囲を決めて、各goroutineでRange GET
    1. 部分ファイルに出力
1. 全部ダウンロードできたら、ファイルをマージして部分ファイル削除

## 工夫した点

- 1つのgoroutineがダウンロードするデータが1MBを超えた場合、1MBごとにファイル出力する。
- Rangeアクセスに対応していなかったら普通にダウンロードする。
- すでにダウンロード済のデータがあった場合(そのパートのファイルがあった場合)、それを再利用する。

## 困っていること

- `go test`のやり方がいまいちわかっていない。依存関係の解決の仕方が理解に至っていない。
  - `go test *.go`だと`export_test.go`が解決できない
  - `go test ./`だと解決できる
  - `go test ./...`すると`utils`以下がimport cycleで失敗してしまう
  - `go test ./utils/*`だと`utils`だと問題ない
    - `go test -cover ./utils/*`でちゃんとカバレッジがどれない
