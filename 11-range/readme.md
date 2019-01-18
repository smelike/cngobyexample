Range

range 用于迭代在各种数据结构中的元素。


使用 range 来统计切片（slice）中所有数字之和。对于 Arrays 的用法一样。


range 作用于 arrays 和 slices，每一次既会有索引(index) 和值(value)。

不需要 index，使用 _ (the blank identifier)忽略。

range 作用于 map，可以获取到 key/value 对。

range 同样可以只获取 map 中的 key。


range 作用于字符串(Unicode code points)。

range on strings iterates over Unicode code points. 

The first value is the starting byte index of the rune and the second the rune itself.

