# Gopher道場#6 課題4

## おみくじAPIを作ってみよう

* JSON形式でおみくじの結果を返す
* 正月（1/1-1/3）だけ大吉にする
* ハンドラのテストを書いてみる

### ビルド

```
$ go build -o omikuji main.go
```

### Usage

```sh
$ omikuji <options>

options
-p [ポート番号] : listenポート番号 (デフォルト:8080)
```

* Ctrl+Cでサーバー終了
* / へのGETアクセスでランダムにおみくじ結果をjsonで返す
* 1/1〜1/3は必ず大吉が返る

----

## テスト実行

```sh
$ go test -v github.com/gopherdojo/dojo6/kadai4/torotake/pkg/omikuji
=== RUN   TestServer_Handler
=== RUN   TestServer_Handler/正月期間のときは全部大吉_開始境界_(1/1_00:00:00)
=== RUN   TestServer_Handler/正月期間のときは全部大吉_終了境界_(1/3_23:59:59.999999999)
=== RUN   TestServer_Handler/正月期間のときは全部大吉_開始境界直前_(12/31_23:59:59.999999999)
=== RUN   TestServer_Handler/正月期間のときは全部大吉_終了境界直後_(1/4_00:00:00)
=== RUN   TestServer_Handler/正月期間以外の時にランダム
--- PASS: TestServer_Handler (0.01s)
    --- PASS: TestServer_Handler/正月期間のときは全部大吉_開始境界_(1/1_00:00:00) (0.00s)
    --- PASS: TestServer_Handler/正月期間のときは全部大吉_終了境界_(1/3_23:59:59.999999999) (0.00s)
    --- PASS: TestServer_Handler/正月期間のときは全部大吉_開始境界直前_(12/31_23:59:59.999999999) (0.00s)
    --- PASS: TestServer_Handler/正月期間のときは全部大吉_終了境界直後_(1/4_00:00:00) (0.00s)
    --- PASS: TestServer_Handler/正月期間以外の時にランダム (0.00s)
PASS
ok      github.com/gopherdojo/dojo6/kadai4/torotake/pkg/omikuji 0.020s
```
