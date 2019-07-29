
### 標準パッケージでどのように使われているか

fmt.Fprintでは以下のように記述されている。
```
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {  
   p := newPrinter()  
   p.doPrintf(format, a)  
   n, err = w.Write(p.buf)  
   p.free()  
   return  
}
```
同じファイル内のFprint内での呼び出しは以下。

```
Fprintf(os.Stdout, format, a...)
```

io.writerのインターフェースはwriteをメソッドに持っている型なら引数に設定できるので、os.Stdoutを引数に指定し、呼び出している。

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

- コードが簡潔にできるために使用することができる。
- interfaceがないと、それぞれの型を引数に指定し、複数の関数を作らなくてはいけない。
- それがinterfaceを設定できることで、interfaceの関数は使用できることが約束される。
- io.Readerの型を満たしているということはReadメソッドが使用できることが約束されており、引数にio.Readerを指定することで関数内でその引数の型の違いを気にすることなく、Readメソッドを使うことができ、シンプルにコードを記述することができる。
