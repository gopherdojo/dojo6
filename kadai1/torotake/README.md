# Gopher道場#6 課題1

## 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください
- [x] mainパッケージと分離する
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
- [x] 準標準パッケージ：golang.org/x以下のパッケージ
- [x] ユーザ定義型を作ってみる
- [x] GoDocを生成してみる

----

## ビルド

```sh
go build imgconv.go
```

## Usage

```sh
imgconv <options> [directory]

options
-i format : input image format "jpg"(default), "png", "gif", "bmp"
-o format : output image format "jpg", "png"(default), "gif", "bmp"
```

----

## 反省
* ユーザー定義型を意味ある形で使えている気がしません
* テストまで書けませんでした…
