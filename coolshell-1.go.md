[Go 语言简介-语法](https://coolshell.cn/articles/8460.html)
[Go 语言简洁-特性](https://coolshell.cn/articles/8489.html)

***注意：这是 2012 年写的文章***
----

> Hello World

```
// hello.go
package main // 申明本文件的 package 名

import "fmt" // import 语言的 fmt 库 - 用于输出

func main() {
  fmt.Println("Hello World")
}
```

> 运行

```
$go run hello.go

$go build hello.go
```
----

> 自己的 package

你可以使用 GOPATH 环境变量，或是使用相对路径来 `import` 自己的 package. （这个规约似乎是否废弃了？@2020）

Go 的规约是：
  - 在 import 中，你可以使用相对路径，如 `./` 或 `../` 来引用你的 packge
  - 如果没有使用相对路径，那么，go 会去找 `$GOPATH/src/` 目录

```
import "./pkg"  // import 当前目录里 pkg 子目录里的所有 go 文件（***似乎不生效）
```

```
import "pkg"  // import 环境变量 $GOPATH/src/pkg 里的所有 go 文件
```

----

> fmt 输出格式

fmt 包与 libc里的那堆使用 `printf`, 'scanf', 'fscanf' 很相似。
注意：`Println` 不支持，`Printf` 才支持 %式的输出：

```
package main

import "fmt"
import "math"

/* 
import (
  "fmt"
  "math"
)

*/


func main() {
  fmt.Println("Hello World, this is fmt.Println")
  
  fmt.Printf("%t\n", 1==2)
  fmt.Printf("二进制：%b\n", 255) // 1111 1111
  fmt.Printf("八进制：%o\n", 255) // 377
  fmt.Printf("十六进制: %X\n", 255) // FF
  fmt.Printf("十进制：%d\n", 255) // 255
  fmt.Printf("浮点数：%f\n", math.Pi) //
  fmt.Printf("字符串：%s\n", "hello world")
}
```

----

> 变量和常量

变量的声明很像 javascript, 使用 `var` 关键字。注意：**go 是静态类型的语言**，下面是代码：

```
// 声明初始化一个 int 变量
var x int = 100
var str string = "hello world"

// 声明初始化多个变量
var i, j, k int = 1, 2, 3

// 不用指明类型，通过初始化值来推导
var b = true // bool 型

```

还有一种定义变量的方式（近似 Pascal 语言）

```
x := 100 // 等价于 var x int = 100
```

常量很简单，使用 `const` 关键字

```
const s string = "hello world"
const pi float32 = 3.1415926
```

----

> 数组

```
func main() {
  var a [5]int
  fmt.Println("array a:", a)
  
  a[1] = 10
  a[3] = 30
  fmt.Println("assign a:", a)
  fmt.Println("len:", len(a))
  
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("b-init:", b)
  
  var c [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      c[i][j] = i + j
    }
  }
  fmt.Println("2d:", c)
}
```

> 数组的切片操作

```
a := [5]int{1, 2, 3, 4, 5}

b := a[2:4] // a[2] 和 a[3]，但不包括 a[4]
fmt.Println(b)

b := a[:4] // 从 a[0] 到 a[4]，但不包括 a[4]
fmt.Println(b)

b := a[2:] // 从 a[2] 到 a[4]，且包括 a[2]
fmt.Println(b)
```

----  

> 分支循环语句

- if 语句

注意: if语句没有圆括号，而必须有花括号

```
// if 语句
if x % 2 == 0 {
  // ...
}

// if-else
if x % 2 == 0 {
  // 偶数
} else {
  // 奇数
}

// 多分支

if num < 0 {
  // negative number
} else if num == 0 {
  // equal to zero
} else {
  // positive number
}

```

- switch 语句

注意：switch 语句没有 break，还可以使用逗号 case 多个值

```
switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  case 4,5,6:
    fmt.Println("four, five, six")
  default:
    fmt.Println("invalid value")
}
```

- for 语句

注意：Go 语言没有 while

```
// init; condition; post
for i := 0; i < 10; i++ {
  fmt.Println(i)
}

//精简的 for 语句 codintion
i := 1
for i<10 {
  fmt.Println(i)
  i++
}

// 死循环的 for 语句 相当于 for(;;)
i := 1
for {
  if i > 10 {
    break
  }
  i++
}
```

----

> 关于分号

Go 使用词法分析器在扫描源代码过程中，使用简单的规则自动插入分号。

规则是这样的：如果在一个新行前方的最后一个标记是一个标识符（包括像 `int` 和 `float64` 这样的单词）、
一个基本的如数值这样的文字、或以下标记中的一个时，会自动插入分号：

```
break continue fallthrough return ++ -- ) }
```
通常 Go 程序仅在 `for` 循环语句中使用分号，以此来分开初始化器、条件和增量单元。如果你在一行中写多个语句，也需要用分号分开。

注意：无论任何时候，你都不应该将一个控制结构(if, for, switch 或 select)的左大括号放在下一行。

----

> map 

map 在别的语言里可能较哈希表或 dict

```
// map 操作代码

func main() {
  m := make(map[string]int) // 使用 make 创建一个空的map
  
  m["one"] = 1
  m["two"] = 2
  m["three"] = 3
  
  fmt.Println(m) // map[three:3, two:2, one:1] 顺序在运行时可能不一样
  fmt.Println(len(m)) // 3
  
  v := m["two"] // 从 mao 里取值
  fmt.Println(v)
  
  delete(m, "two") // 删除
  fmt.Println(m)
  
  m1 := map[string]int{"one": 1, "two": 2, "three": 3}
  fmt.Println(m1) // 顺序在运行时可能不一样
  
  // 使用 range 遍历 map
  for key, val := range m1 {
    fmt.Println("%s => %d \n", key, val)
  }
}
```

> 注意: map 中的index 不允许使用单引号。

----

> 指针

```
var i int = 1
var pInt *int = &i

fmt.Println("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

*pInt = 2
fmt.Println("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

i = 3
fmt.Println("i=%d\tpInt=%p\t*pInt=%d", i, pInt, *pInt)
// 3

----

> 内存分配

Go 具有连个分配内存的机制，分别是内建的函数 `new`  和 `make`。他们所做的事不同，所应用到的类型也不同，这可能引起混淆，但规则却很简单。

new 是一个分配内存的内建函数，但不同于其他语言中同名的 new 所作的工作，**它只是将内存清零，而不是初始化内存**。

new(T)为一个类型为 T 的新项目分配了值为零的存储空间并返回其地址地，也就是一个类型为 `*T`的值。

用 Go 的术语来说，就是它返回了一个指向分配的类型为 T 的零值的指针。


`make(T, args)` 函数的目的与 `new(T)` 不同。它仅用于创建切片、map 和 chan(消息管道)，并返回类型 `T` (不是 *T) 的一个被初始化了的（不是零）实例。

这种差别的出现是由于这三种类型实质上是对在使用前必须进行初始化的数据结构的引用。

例如，切片是一个具有三项内容的描述符，包括指向数据（在一个数组内部）的指针、长度以及容量，在这三内容被初始化之前，切片值为 `nil`。

对于切片、映射和信道，`make` 初始化了其内部的数据结构并准备了将要使用的值。如：

下面的代码分配了一个整型数组，长度为 10，容量为 100，并返回前 10 个数组的切片

```
make([]int, 10, 100)
```

以下示例说明了 `new` 和 `make` 的不同。

```
var p *[]int = new([]int) // 为切片结构分配内存；*p == nil; 很少使用
var v []int = make([]int, 10) // 切片 v 现在是对一个新的有 10 个整数的数组的引用

// 不必要地使问题复杂化
var p *[]int = new([]int)
fmt.Println(p) // &[]
*p = make([]int, 10, 10)
fmt.Println(p) // &[0 0 0 0 0 0 0 0 0 0]
fmt.Println((*p)[2]) // 0

// 习惯用法
v := make([]int, 10)
fmt.Println(v) // []

```
