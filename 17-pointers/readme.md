Pointers

Go 支持指针，允许通过传值和记录的引用。(pass references to values and records)

通过函数 zeroval 和 zeroptr 对比，指针和值的使用。

zeroval 有着一个 int 参数，参数会被以值传递。

zeroval 将会获取 ival 的复制，不同于调用函数内的 ival。


作对比的 zeroptr 函数，有着一个 *int 参数，该形式代表接收一个 int 指针参数。

函数体内的代码 *iptr，通过将自己的内存地址重新指向到当前值的内存地址。赋值给它(dereferenced pointer)，改变所指向内存地址的保存的值。
(dereferences the pointer from its memory address to the current value at that address)


&i 语法，代表 i 的内存地址，如：i 指针。

指针同样可以被打印。

zeroval 不改变 i 。

zeroptr 改变 i，因为 zeroptr 已关联到该变量的内存地址。




