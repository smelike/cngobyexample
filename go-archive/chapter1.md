Go 是一门编译型语言。

Go 语言的工具链将源代码及其依赖转换成计算机的机器指令。

Go 语言提供的工具都通过一个单独的命令 go 调用，go 命令有一系列子命令。

最简单的一个子命令就是 run。这个命令编译一个或多个以 .go 结尾的源文件，链接库文件，并运行最终生成的可执行文件。


---

```
// helloworld.go
import main
import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

> go run helloword.go

Go 语言原生支持 Unicode，可以处理全世界任何语言的文本。

---

如果你希望能够编译这个程序，保存编译结果以备将来只用。可以用 build 子命令：

> go build helloworld.go

这个命令生成一个名为 helloworld的可执行的二进制文件，之后你可以随时运行它，不需任何处理。

> ./helloworld # Linux/Mac OS

> ./helloworld.exe




