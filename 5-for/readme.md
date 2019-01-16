for 是 Go 的唯一循环结构。

三种基础的 for 循环类型：

最简单的是，只有有一个条件。

```
	i := 1
	
	for i < 3 {
	
	} 

```

典型的是，initial/condition/after for 结构。


一个条件都没有的 for 结构，会一致循环直到使用 break 中断循环，或者出现 return(return from the enclosing function)。

还可使用 continue 跳过当前循环，进入下一次循环。


我们可以看到其他的 for 形式，当我们学习到 range 声明，channels，与其他数据结构。
