
Maps

Maps 是 Go 内置的关联数据类型(associative data type)。
（其他开发语言称之为 hashes or dicts）


创建一个空 map，使用内置 make: make(map[key-type]val-type)。

设置 key/value 对，使用 name[key] = val 语法。

打印 map，使用 fmt.Println 可将所有 key/value 对打印出来。


给一个 key 赋值，使用 name[key]。


对 map 使用 len，返回 key/value 对的数量。

delete 可以移除 map 的 key/value 对。


从 map 获取一个值时，有一个可选的返回值，该返回值代表是否该 key 存在于 map。

该【可选返回值】，可用于区分不存在的 key(missing keys)，与 key 对应的值为 0 或 ""。

不需要 key 对应的值，可以使用“the blank identifier _”。


在一行内声明和初始化一个新的 map，使用同样的语法：n := map[string]int{"foo": 1, "bar": 2}


注意：使用 fmt.Println 打印 map，输出格式为：map[k:v k:v]。


