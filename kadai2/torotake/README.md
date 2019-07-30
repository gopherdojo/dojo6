# Gopher道場#6 課題2

## io.Readerとio.Writerについて調べてみよう

### 標準パッケージでどのように使われているか

* 標準入力、標準出力、標準エラー出力 (osパッケージ)
* 暗号や圧縮などデータ変換ソース (crypt, gzipパッケージなど)
* ネットワーク通信の送受信 (netパッケージなど)
* 各種Reader/WriterをラッピングしてI/O処理にバッファリング機能を追加 (bufioパッケージ)


### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

ファイル、標準入出力、ネットワーク通信など様々なIO処理について、同じ取り扱う事が出来、入出力先の切り替えも容易となる。

例えば、fmt.Fprintfでは出力対象としてio.Writerを要求するので渡すものを変えるだけでファイルへの書き込み、標準出力への出力、ネットワークソケットへの送信処理等を行う事が可能。

```go
// 標準出力
fmt.Fprintf(os.Stderr, "Hello World.")

// ファイルへ書き込み
fp := os.Open("hoge.txt")
fmt.Fprintf(fp, "Hello World.")

// ネットワークへ送信
conn, _ := net.Dial("tcp", "example.com:80")
fmt.Fprintf(conn, "Hello World.")
```

----

## 1回目の宿題のテストを作ってみて下さい

- [ ] テストのしやすさを考えてリファクタリングしてみる
- [x] テストのカバレッジを取ってみる
- [x] テーブル駆動テストを行う
- [ ] テストヘルパーを作ってみる

テスト実行
```
$ go test -v github.com/gopherdojo/dojo6/kadai2/torotake/convert

=== RUN   TestConvert
=== RUN   TestConvert/ConvertJpegToPng
convert testdata/test/test_jpg.jpg -> testdata/test/test_jpg.png ...
=== RUN   TestConvert/ConvertJpegToBmp
convert testdata/test/test_jpg.jpg -> testdata/test/test_jpg.bmp ...
=== RUN   TestConvert/ConvertJpegToGif
convert testdata/test/test_jpg.jpg -> testdata/test/test_jpg.gif ...
=== RUN   TestConvert/ConvertPngToJpeg
convert testdata/test/test_png.png -> testdata/test/test_png.jpg ...
=== RUN   TestConvert/ConvertPngToBmp
convert testdata/test/test_png.png -> testdata/test/test_png.bmp ...
=== RUN   TestConvert/ConvertPngToGif
convert testdata/test/test_png.png -> testdata/test/test_png.gif ...
--- PASS: TestConvert (1.77s)
    --- PASS: TestConvert/ConvertJpegToPng (0.32s)
    --- PASS: TestConvert/ConvertJpegToBmp (0.04s)
    --- PASS: TestConvert/ConvertJpegToGif (0.70s)
    --- PASS: TestConvert/ConvertPngToJpeg (0.06s)
    --- PASS: TestConvert/ConvertPngToBmp (0.03s)
    --- PASS: TestConvert/ConvertPngToGif (0.59s)
=== RUN   TestListFiles
--- PASS: TestListFiles (0.00s)
PASS
ok      github.com/gopherdojo/dojo6/kadai2/torotake/convert     1.776s
```

カバレッジ
```
$ go test -coverprofile=profile github.com/gopherdojo/dojo6/kadai2/torotake/convert

ok      github.com/gopherdojo/dojo6/kadai2/torotake/convert     1.771s  coverage: 77.8% of statements
```
