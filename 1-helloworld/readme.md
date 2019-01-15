第一个程序是打印"hello world"。


创建文件 hello-world.go

代码如下：

```
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

```


如何运行 hello-world.go 文件？

第一种方式： go run hello-world.go


第二种方式：go build hello-world.go，执行生成的程序，调用 hello-world.exe 即可。


go run 执行运行程序，不成二进制软件包。

go build 生成二进制软件包。Windows 中，生成 .exe 包。*nix 则不同。






