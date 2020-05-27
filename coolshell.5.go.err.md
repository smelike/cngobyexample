> 错误处理 - Panic/Recover

对于***【不可恢复的错误】**，Go 提供了一个内建的 panic 函数，它将创建一个运行时错误并使程序停止（相当暴力）。
该函数接收一个任意类型（往往是字符串）作为程序死亡时要打印的东西。
当编译器在函数的结尾处检查到一个 panic 时，就会停止进行常规的 return 语句检查。

下面的仅仅是一个示例。实际的库函数应避免 `panic`。如果问题可以容忍，最好是让事情继续下去而不是终止整个程序。

```
var user = os.Getenv("USER")

func init() {
  if user == "" {
    panic("no value for $USER")
  }
}
```

当 panic 被调用时，它将立即停止当前函数的执行并开始逐级解开函数堆栈，同时运行所有被 `defer` 的函数。

如果这种解开达到堆栈的顶端，程序就死亡了。但是，也可以使用内建的 `recover` 函数来重新获得 Go 程序的控制并恢复正常的执行。

***对 recover 的调用会通知解开堆栈，并返回传递到 panic 的参量。***

由于仅在解开期间运行的代码处在被 defer 的函数之内，recover 仅在被延期的函数内部才是有用的。

下面是一个例程:

```
func g(i int) {
  if i > 1 {
    fmt.Println("Panic!")
    // panic 不可恢复的错误
    panic(fmt.Sprintf("%v", i))
  }
}

func f() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in f", r)
    }
  }()
  
  for i := 0; i < 4; i++ {
    fmt.Println("Calling g with ", i)
    g(i) // 调用 g 函数
    fmt.Println("Returned normally from g.")
  }
}

func main() {
  f()
  fmt.Println("Returned normally from f.")
}
```





