声明一个变量有好几种方式，下面这些都等价：


```
s := ""
var s string
var s = ""
var s string = ""

```

用哪种不用哪种，为什么呢？

第一种形式，是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。

第二种形式依赖于字符串的默认初始化零值机制，被初始化为 ""。

第三种形式用得很少，除非同时声明多个变量。

第四种形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。

实践中一般使用前两种形式中的某个，初始值重要的话就显式指定变量的类型，否则使用隐式初始化。



