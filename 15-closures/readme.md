Closure

Go 支持匿名函数(anonymous functions)，可以作为闭包(closures)。

匿名函数，在你想在一行内定义一个函数，而不想命名时。

函数 intSeq 返回一个函数，该返回函数是在 intSeq 函数体内定义的匿名函数。

```
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

```

该返回函数，通过返回变量 i 作为结束，形成一个闭包。

调用函数 intSeq ，将返回值赋给 nextInt。

nextInt 函数会保存自己的 i 值，每次被调用都会更新。

函数的另一个特性是：递归(recursion)。


