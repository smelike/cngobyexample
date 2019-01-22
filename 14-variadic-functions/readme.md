
Variadic Functions 

可变函数


可变函数，在调用时，可以使用任意数量的尾随参数(trailing arguments)。

如，fmt.Println 就是一个常见的可变函数。


定义一个可接收任意数量的 ints 参数的函数。

```
func sum(nums ...int) {
	// ...
}

```

可变函数，亦可如普通函数一样，通过参数传值调用。


如果在一个切片中有多个 args，可以使用到可变函数中，如： func(slice...) 这样。


Go 函数的另一个关键是闭包。

