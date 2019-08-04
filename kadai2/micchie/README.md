# Gopher 道場
## 6-2-1: io.Readerとio.Writer
io.Reader と io.Writer について調べてみよう。

### 1. 標準パッケージでどのように使われているか
- バッファ付入出力:
	- https://golang.org/pkg/bufio/
- 文字列フォーマット:
	https://golang.org/pkg/fmt/
		- Fprint, Fprintf, Fprintln
		- Fscan, Fscanf, Fscanln
- JSON エンコード・デコード:
	- https://golang.org/pkg/encoding/json/
		- NewDecorder, NewEncoder
- HTTP:
	- https://golang.org/pkg/net/http/

### 2. io.Reader と io.Writer があることでどういう利点があるのか具体例を挙げて考えてみる
- 入出力を一律 io.Reader, io.Writer のインターフェースで抽象化しているため, 「読み書きをする」ことに対してはすべて共通して扱える。
- io.Pipe を利用することで, 複数の Read, Write を効率よくつなげて処理できるので, 大量のデータを扱う場合はメモリの節約ができるのでは。
	- https://golang.org/pkg/io/#Pipe

## 6-2-2: テストを書いてみよう
1回目の宿題のテストを作ってみて下さい。

- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる
