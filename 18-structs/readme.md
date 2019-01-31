Structs

Go's structs are typed collections of fields.

structs 是字段类型的集合。

They're useful for grouping data together to form records.

structs 用于集合数据称为记录非常合适。

This person struct type has name and age fields.

```
type person struct {
	name string
	age int
}

```

This syntax creates a new struct.

创建一个新的 struct，语法：

```
person{"Bob", 20}
```

You can name the fields when initializing a struct.
struct 初始化，可以命名字段。

```
person{name: "Alice", age: 30}
```

Omitted fields will be zero-valued.

省略的字段，会被赋予零值。

An & prefix yields a pointer to the struct.

& 产生一个 struct 指针

Access struct fields with dot.

. 可以访问 struct 字段。


You can also use dots with struct pointers - the pointers are automatically dereferenced.

通过 struct 指针，一样可以使用 .

Structs are mutable.

structs 是可变的。
