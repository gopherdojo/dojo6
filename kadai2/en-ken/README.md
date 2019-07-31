# 課題2

## io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

### 標準パッケージで使われているところ

- 基本的にI/Oが存在するところのI/Fには存在する。
  - stdin, stdout, stderr (os.Stdin, Stdout, Stderr)
  - ファイル(os.File)
  - メモリ(bytes.Buffer)
  - TCP,UDP (net)
  - POSTのbody (net.http)
- bufioではラップしてより使いやすくしている。

### 利点

- 理論的にはio.Reader/WriterのI/Fを実装しているどんなシステムにでも着替えられる。
  - 入力出力を抽象化できるので、呼び出し側は極論I/Oのシステムが何か意識しなくて良い。
  - 呼び出し側の実装が外界に依存しない。
  - DIPやりやすい。
  - 粗結合。移植観点からも有利。

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

## カバレッジ結果

```bash
> ./cov.sh
ok      github.com/gopherdojo/dojo6/kadai2/en-ken/cli   0.001s  coverage: 74.1% of statements
ok      github.com/gopherdojo/dojo6/kadai2/en-ken/imgcnv        0.071s  coverage: 85.0% of statements
ok      github.com/gopherdojo/dojo6/kadai2/en-ken/kadai2        1.150s  coverage: 0.0% of statements
```

## kadai1からの変更点

- mainをkadai2のディレクトリに変更
- 冗長だった構造体のI/Fから関数のI/Fに変更
  - 関数のI/Fよき
- コマンド引数をシンプルに変更
- 独自モックを実装
- 独自実装していたファイル取得の再起処理を`filepath.Walk`に変更
