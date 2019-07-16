## 概要
指定ディレクトリ配下にある画像ファイルを、再帰的に走査して別の形式のファイルに変換します。png, jpeg, gifに対応しています。

## 実行方法

```
go run main.go -in TARGET_PATH -src SOURCE_EXTENSION -dst DESTINATION_EXTENSION

-in 走査したいディレクトリを指定します。
-src 変換元のファイル形式を指定します。
-dst 変換後のファイル形式を指定します。
```