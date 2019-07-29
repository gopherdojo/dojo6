## 概要
指定ディレクトリ配下にある画像ファイルを、再帰的に走査して別の形式のファイルに変換します。
png, jpeg, gifに対応しています。

## 実行方法

```
imgconv -in INPUT_PATH -out OUTPUT_PATH -src SOURCE_EXTENSION -dst DESTINATION_EXTENSION

-in 走査したいディレクトリを指定します。
-out 出力先のディレクトリを指定します。
-src 変換元のファイル形式を指定します。
-dst 変換後のファイル形式を指定します。
```

## テストカバレッジ
```
$ go test -cover imgconv/...
ok      imgconv 0.434s  coverage: 88.2% of statements
ok      imgconv/converter       0.430s  coverage: 100.0% of statements
```

### io.ReaderとioWriterについて
```
type Reader interface {
        Read(p []byte) (n int, err error)
}
```
byte型のスライスを引数に取り、pに値を読み込む。
読み込んだバイト数とエラーを返す。

```
type Writer interface {
        Write(p []byte) (n int, err error)
}
```
byte型のスライスを引数に取り、pから値を書き込む。
書き込んだバイト数とエラーを返す。

- 標準パッケージでどのように使われているか
	- tarやzip等のアーカイブ形式
	- bzip2やlzw等の圧縮形式
	- ECDSAやtls等の暗号化形式
	- base64や16進数等のフォーマット
	- gifやpng等の画像形式
	- 標準入出力や標準エラー出力
	- http通信でのリクレス等
	- MIMEマルチパートやQuoted-printable等の方式
	- ファイル
	- 文字列やバイトバッファ
	- ログ

上記のように様々な箇所での読み書きに使用されていた。
また、以下のようなerrorの戻り値を利用するような使い方も見られた。
```
type errorReader struct {
	error
}
func (r errorReader) Read(p []byte) (n int, err error) {
	return 0, r.error
}
```

```
type eofReader struct{}
func (eofReader) Read([]byte) (int, error) {
	return 0, EOF
}
```

- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
	- io.Readerやio.Writerを実装したインスタンスをインターフェースを介して利用することで、使用する側は実体が何かを意識せずに扱うことができる。
	- シンプルなインターフェースなので実装が容易。何に対して読み書きを行うかは定義されていないので、入出力先に応じて適宜実装すれば良い。
	- テストやデバッグ時に、実際の処理とは別の処理に置き換えるといった差し替えが容易に行える。
	- ioパッケージのpipeやio/ioutilパッケージで定義された既存の便利メソッドを利用できる。