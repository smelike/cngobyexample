
查找重复的行

对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：

一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。

```
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		counts[input.Text()]++
	}
	
}
```

正如 for 循环一样，if 语句条件两边也不加括号，但是主题部分需要加。if 语句的 else 部分是可选的，在 if 的条件为 false 时执行。

map 存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。
键可以时任意类型，只要其值能用 == 运算符比较，最常见的例子是字符串；值则可以是任意类型。

内置函数 make 创建空 map，此外它还有别的作用。

（从功能和实现上说，Go 的 map 类似于 Java 语言中 HashMap，Python 语言中的 dict，Lua 语言中的 table，通常使用 hash 实现。）

每次 dup 读取一行输入，该行被当做 map，其对应的值递增。counts[input.Text()]++语句等价下面两句：

```
line := input.Tex()
counts[line] = counts[line] + 1

```
map 中不含某个键时不用担心，首次读到新行时，等号右边的表达式 counts[line] 的值将被计算为其类型的零值，对于 int 即 0。


bufio 包，它使处理输入和输出方便又高效。Scanner 类型是该包最有用的特性之一，它读取输入并将其拆成并将其拆成行或单词；通常是处理形式的输入最简单的方法。


程序使用短变量声明创建 bufio.Scanner 类型的变量 input。


> input := bufio.NewScanner(os.Stdin)

每次调用 input.scanner，即读入下一行，并移除行末的换行符；读取的内容可以调用 input.Text() 得到。
Scan 函数在读到一行时返回 true，在无输入时返回 false。

fmt.Printf 函数对一些表达式产生格式化输出。该函数的首个参数是歌格式字符串，指定后续参数该如何格式化。
各个参数的格式取决于“转换字符”（conversion character），形式为百分号后跟一个字母。举个例子，%d 标识以十进制形式打印一个整型操作数，而 %s 则标识把字符串型操作数的值展开。


%d 十进制整数

%x, %o, %b 十六进制，八进制，二进制整数

%f, %g, %e 浮点数：3.14.1593 3.141592653589793 3.141593e+00
%t 布尔：true 或 false
%c 字符（rune）（Unicode码点）
%s 字符串
%q 带双引号的字符串"abc"或带单引号的字符'c'
%v 变量的自然形式(natural format)
%T 变量的类型
%% 字面上的百分号标志（无操作数）

默认情况下，Printf 不会换行。
按照惯例，以字母 f 结尾的格式化函数，如 log.Printf 和 fmt.Errorf，都采用 fmt.Printf 的格式化准则。
而以 ln 结尾的格式化函数，则遵循 Println 的方式，以跟 %v 产不多的方式格式化参数，并在最后添加一个换行符。
（f 指 format，ln 指 line。）

---

os.Open 函数返回两个值。

第一个值是被打开的文件（*os.File），其后被 Scanner 读取。

os.Open 返回的第二个值是内置 error 类型的值。如果 err 等于内置值 nil（相当于其他语言里的 NULL），那么文件被成功打开。
读取文件，直到结束，然后调用 Close 关闭该文件，并释放占用的所有资源。

相反的话，如果 err 的值不是 nil，说明打开文件时出错了。这种情况下，错误描述了所遇到的问题。我们的错误处理非常简单，只是使用
Fprintf 与表示任意类型默认格式值得动词%v，向标准错误流打印一条信息了，然后dup继续处理下一个文件；
continue 语句直接跳到 for 循环的下个迭代开始执行。


为了使示例代码保持合理的大小，有意简化了错误处理，应该检查 os.Open 返回的错误值，然而，使用 input.Scan 读取文件过程中，不大可能出现错误。

注意 countLines 函数在其声明前被调用。函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用。

map 是一个由 make 函数创建的数据结构的引用。
map 作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对 map 底层数据结构的任何修改，调用者函数都可以通过持有的 map 引用看到。

在我们的例子中，countLine 函数向 counts 插入的值，也会被 main 函数看到。（类似于 C++ 里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）


dup 的前两个版本以“流”模式读取输入，并根据需要拆分成多个行。
理论上，这些程序可以处理任意数量的输入数据。还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。

dup3 就是这么操作的。这个例子引入了 ReadFile 函数（来自 io/ioutil 包），其读取指定文件的全部内容，strings.Split 函数把字符串分割成子串的切片。（ReadFile 只读制定文件，不读标准输入）



