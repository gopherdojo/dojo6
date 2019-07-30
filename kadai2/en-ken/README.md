# 課題2

## io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

### 標準パッケージで使われているところ

### 利点

## 1回目の宿題のテストを作ってみて下さい

- [x] テストのしやすさを考えてリファクタリングしてみる
- [x] テストのカバレッジを取ってみる
- [x] テーブル駆動テストを行う
- [x] テストヘルパーを作ってみる

## 使い方

```go
kadai2 -in [INPUT_EXT] -out [OUTPUT_EXT] SRC_DIR DST_DIR

#Usage of kadai2:
#  -in string
#        input extension (jpg/png) (default "jpg")
#  -out string
#        output extension (jpg/png) (default "png")
```

## kadai1からの変更点

- mainをkadai2のディレクトリに変更
- 冗長だった構造体のI/Fから関数のI/Fに変更
- コマンド引数をシンプルに変更
- 独自モックを実装
