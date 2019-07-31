# 課題 2

## io.Readerとio.Writer
io.Readerとio.Writerについて調べてみよう

### 標準パッケージではどのように使われているか
私がよく使ったのは`func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`です。
第一引数が`io.Writer`となっているため、条件を満たしているものは何でも引数に指定できるため頻繁に使っています。
`io.Writer`は`Write(p []byte) (n int, err error)`メソッドを持っていればよい。
このため、`func (f *File) Write(b []byte) (n int, err error)`を持つ`os.File`オブジェクトを指定すればファイルに出力することができる。

io.Readerは競技プログラミングでお世話になります。
読み込み高速化のために`bufio.NewScanner`を頻繁に使いますが、これは`func NewScanner(r io.Reader) *Scanner`です。
僕は`bufio.NewScanner(os.Stdin)`をよく使いますが、`os.Stdin`は、`Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")`です。
NewFile関数について調べると`func NewFile(fd uintptr, name string) *File`であり、Fileオブジェクトを返します。
Fileオブジェクトは`func (f *File) Read(b []byte) (n int, err error)`であるため、`io.Reader`を満たします。
普段は、`os.Stdin`しか使いませんが、別の入力を読み込む時は`buio.NewScanner`の引数を`io.Reader`を満たしているものに変えれば同様に読み込めます。

### io.Readerとio.Writerがあることでどういう利点があるのかを具体的に挙げて考えてみる。

上記でも、触れた通り書き込むこと、読み込むことに関して抽象化されていることによりプログラマーがほとんど意識を割く必要がない点に大きな利点があると考えます。

## テストを書いてみよう
1回目の宿題のテストを作ってみてください

### テストのしやすさを考えてリファクタリングしてみる

### テストのカバレッジを取ってみる

### テーブル駆動テストを行う

### テストヘルパーを作ってみる

## Install

```
export GOBIN=`pwd`/_bin
$ go install github.com/gopherdojo/dojo6/kadai2/dobuzora/cmd/j2p
$ _bin/cmd
```

## How to use

```
./_bin/j2p DIRPATH
```

## 回答
