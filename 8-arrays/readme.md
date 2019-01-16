Go 中，数组是一个有着具体长度的、有数字编号的元素组合。

创建一个数组存放 5 个整数。

元素的类型和长度皆属于数组的部分类型(the array's type)。


数组操作


对一个索引(index)上 进行赋值，语法；array[index] = value

使用索引(index)，获取一个值(value)，语法：array[index]


数组长度

使用内置函数 len，可获取数组的长度。


声明并初始化一个数组，如： b := [5]int{1, 2, 3, 4, 5}


数组类型(Array types) 都是一维的，你可使用多个类型来构建多维度的数据结构，如下：

> var towDimensionalData [2][3]int


打印数组

使用 fmt.Println 打印输出数组，结果格式是：[v1 v2 v3 ...]


更多，我们会在程序中使用数组切片，而不是整个数组。详情见于下一个知识点：Slices。
