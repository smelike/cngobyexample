> 错误处理 -Defer

C 语言处理，有资源泄露的问题

```
func CopyFile(dstName, srcName string) (written int64, err error) {
  src, err := os.Open(srcName)
  if err != nil {
    return 
  }
  
  dst, err := os.Create(dstName)
  if err != nil {
    return 
  }
  written, err = io.Copy(dst, src)
  dst.Close()
  src.Close()
  return
}
```

Go 语言引入了 Defer 来确保那些被打开的文件能被关闭。如下所示：

```
func CopyFile(dstName, srcName string) (written int64, err error) {
  src, err := os.Open(srcName)
  if err != nil {
    return 
  }
  defer src.Close()
  
  dst, err := os.Create(dstName)
  if err != nil {
    return 
  }
  defer dst.Close()
  
  return io.Copy(dst, src)
}

```

Go 的 defer 语句预设一个函数调用（**延期的函数**），该调用在函数执行 defer 返回时立刻运行。
该方法显得不同常规，但是处理上述情况很有效，无论函数怎样返回，都必须进行资源释放。

再来看一个 defer 函数的示例：

```
for i := 0; i < 5; i++ {
  defer fmt.Println("%d", i)
}
```

被延期的函数以【后进先出】的顺序执行，因此上代码返回结果时将打印 4 3 2 1 0。

（defer 的函数行为，对 defer 的理解还是有点模糊）
