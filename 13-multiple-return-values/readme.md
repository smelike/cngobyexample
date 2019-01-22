
Multiple Return Values


Go 支持多返回值，该特性常被使用，例如：返回一个函数 result 和 error values。


函数返回两个整数值，(int, int) ，如: 

```
func vals() (int, int) {

	return 3, 7
}

```

通过多个参数赋值，来使用两个返回值，如：

```
a, b := vals()

```

只使用其中某个返回值，使用 _ 忽略不需要的返回值，如：

```

_, c := vals()

```
