使用 Go 语言标准库里的 image 包，生成一系列的 bit-mapped 图，容纳后将这些图片编码为一个 GIF 动画。

代码用了一些新的结构，包括 const 声明，struct 结构体类型，复合声明。

> var palette = []color{color.White, color.Black}

```
const (
	whiteIndex = 0
	blackIndex = 1
)

```

---

程序里的常量声明给出了一系列的常量值，常量是指在程序编译后运行时始终都不会变化的值，比如圈数、帧数、延迟值。

常量声明和变量声明一般都会出现在包级别，所以这些常量在整个包中都是可以共享的，或者你也可以把常量声明定义在函数体内部，那么这种常量就只能在函数体内用。

目前常量声明的值必须是一个数字值、字符串或者一个固定的 boolean 值。


```
// 生成一个切片
var palette = []color{color.White, color.Black}

// 生成一个结构体
gif.GIF{LoopCount: nframes}
```
[]color.Color{...} 和 gif.GIF{...} 这两个表达式就是我们说的复合声明。
这是实例化 Go 语言里的复合类型的一种写法。

gif.GIF 是一个 struct 类型。struct 是一组值或者叫字段的集合，不同的类型集合在一个 struct 可以让我们以一个统一的处理。

anim 是一个 gif.GIF 类型的 struct 变量。这种写法会生成一个 struct 变量，并且其内部变量 LoopCount 字段会被设置为 nframes；

而其他的字段会被设置为各自类型默认的零值。struct 内部的变量可以以一个点（.）来进行访问，就像在最后两个赋值语句中显式地更新了 anim 这个struct 的 Delay 和 Image 字段。

lissajous 函数内部有两层嵌套的 for 循环。外层循环会循环 64 次，每一次都会生成一个单独的动画帧。它生成了一个包含两种颜色的 201 & 201 大小的图片，白色和黑色。


---

练习：

1、修改前面的 Lissajous 程序里的调色板，由黑色改为绿色。我们可以用 color.RGBA{0xRR, 0xGG, 0xff} 来得到 #RRGGBB 这个色值，三个十六进制的字符串分别代表红、绿、蓝像素。

2、修改 Lissajous 程序，修改其调色板来生成更丰富的颜色，然后修改 SetColorIndex 的第三个参数，看看现实结果吧。

