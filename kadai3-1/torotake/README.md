# Gopher道場#6 課題3-1

## タイピングゲームを作ろう

* 標準出力に英単語を出す（出すものは自由）
* 標準入力から1行受け取る
* 制限時間内に何問解けたか表示する

### ビルド

```
$ go build -o game main.go
```

### Usage

```sh
game <options>

options
-f [問題文ファイルのパス] : 問題文ファイル (デフォルト:カレントディレクトリのwords.txt)
-t [制限時間(秒)] : 制限時間 (デフォルト:30秒)
```

#### 問題文ファイルについて

* 英単語を並べたテキストファイル
* 1行が1問となる
* 前後の空白はトリムされ、空行は無視される
* 問題はランダムに選択される (重複あり)

----

## テスト実行

```sh
$ go test -v github.com/gopherdojo/dojo6/kadai3-1/torotake/typing
=== RUN   TestGame_Run
--- PASS: TestGame_Run (5.00s)
PASS
ok      github.com/gopherdojo/dojo6/kadai3-1/torotake/typing    5.006s
```
