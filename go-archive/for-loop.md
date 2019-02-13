Go 语言只有 for 循环这一种循环语句。for 循环有多重形式，其中一种如下所示：


for initialization; condition; post {
	// zero or more statements
}

for 循环三个部分不需括号包围。大括号强制要求，左大括号必须和 post 语句在同一行。

initialization 语句是可选的，在循环开始前执行。

initialization 如果存在，必须是一条简单语句（simple statement），即，短变量声明、自增语句、赋值语句或函数调用。

condition 是一个布尔表达式（boolean expression），其值再每次循环迭代开始时计算。

如果为 true 则执行循环体语句。post 语句在循环体执行结束后执行，之后再次对 condition 求值。

condition 值为 false 时，循环结束。

for 循环的这三个部分每个都可以省略，如果省略 initialization 和 post，分号也可以省略：

```
	// a traditional "while" loop
	for condition {
		// ...
	}

```

如果连 condition 也省略了，像下面这样：

```
	// a traditional infinite loop
	
	for {
		//...
	}

```

这就变成一个无限循环，尽管如此，还可以用其他方式终止循环，如一条 break 或 return 语句。

for 循环的另一种形式，在某种数据类型的区间（range）上遍历，如字符串或切片。


---

每次循环迭代，range 产生一对值；索引以及在该索引处的元素值。这个例子不需要索引，但 range 的语法要求，要处理元素，必须要处理索引。

一种思路是把索引赋值给一个临时变量，如 temp，然后忽略它的值，但 Go 语言不允许使用无用的局部变量（local variables），因为这会导致编译错误。


Go 语言中这种情况的解决办法是用“空标识符”（blank identifier），即 _（也就是下划线）。

空标识符可用于任何语法需要变量名但程序逻辑不需要的时候，例如，在循环里，丢弃不需要的循环索引，保留元素值。

大多数的 Go 程序员都会像上面这样使用 range 和 _ 写 echo 程序，因为隐式地而非显式地索引 os.Args，容易写对。


echo 的这个版本使用一条短变量声明来声明并初始化 s 和 seps，也可以将这两个变量分开声明。

如前文所述，每次循环迭代字符串 s 的内容都会更新。+= 连接原字符串、空格和下个参数，产生新字符串，并把它赋值给 s。
s 原来的内容已经不再使用，将再适当时机对它进行垃圾回收。

如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用 strings 包的 join 函数：

```
	
	func main() {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
```









