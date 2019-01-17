
Slices

Go中，切片是一个关键数据类型，提供了比数组(arrays)更强大的排序(sequences)接口。


不同于数组，切片是由所包含元素决定的（而不是元素的数量）。


创建一个长度非零的空切片，可以使用内置函数 make。


案例：创建一个长度为 3 的字符串切片（初始化为默认值）。

```
s := make([]string, 3)
fmt.Println("emp:", s)

```

我们可以像操作数组一样，进行设置(set)和获取(get)操作。


len 一样可以获取切片的长度。


除了这些基本操作之外，切片支持的操作比数组的操作要多。

其中一个是内置的 append 操作，返回包含由一个或多个新值的切片。

注意：必须接收 append操作的返回值，当我们想要一个新的切片。


切片还可以被复制(copy'd)。

创建一个空切片 c，长度与 s 相同，并且从 s 中复制到 c 中。


切片支持“切片”操作，语法：slice[low:high]。

操作案例：从元素中获取一个切片，s[2], s[3] 和 s[4]。


s[5] 获取一个不包括元素 5 的切片。


在一行中声明和初始一个切片变量，如：t := []string{"g", "h", "i"}


切片可以是多维的数据结构。切片内的切片长度，可以不同，不像多维度的数组。


注意：切片是不同于数组的类型，它们都可以使用 fmt.Println 打印输出。


---

https://blog.golang.org/go-slices-usage-and-internals


## Arrays

An array literal can be specified like so:
> b := [2]string{"Penn", "Teller"}

Or, you can have the compiler count the array elements for you:
> b := [...]string{"Penn", "Teller"}

In both cases, the type of b is [2]string.

---

## Slices

The type specification for a slice is []T, where T is the type of the elements of the slice.
Unlike an array type, a slice type has no specified length.


A slice literal is declared just like an array literal, except you leave out the element count:
> letters := []string{"a", "b", "c", "d"}

A slice can be created with the built-in function called make, which has the signature,
> func make([]T, len, cap) []T

where I stands for the element type of the slice to be created. The **make** function takes a type, a length, and an optional capacity.
When called, **make** allocated an array and returns a slice refers to that array.

```
var s []byte
s = make([]byte, 5, 5)
// s = []byte{0, 0, 0, 0, 0}

```

When the capacity argument is omitted, it defaults to the specified length. Here's a more succinct version of the same code:

> s := make([]byte, 5)

The length and capacity of a clie can be inspected using the built-in len and cap functions.

```
len(s) = 5
cap(s) = 5
```

Slice internals

A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity(the maximum length of the segment).


 ptr
*Elem
-----
 len
int
-----
 cap
int
-----

The length is the number of elements referred to by the slice.
The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).
The distinction between length and capacity will made clear as we walk through the next few examples.


Slices does not copy the slice's data. It creates a new slice value that points to the original array.

This makes slice operations as efficient as manipulating array indices. Therefor, modifying the elements (not the slice itself) of a re-slice
modifies the elements of the original slice:

```
	d := []byte{'r', 'o', 'a', 'b'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

```

Earlier we sliced s to a length shorter than its capacity. We can grow s to its capacity by slicing it again:

> s = s[:cap(s)]

A slice cannot be grown beyonf its capacity. Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a slice or array.

Similarly, slices cannot be re-sliced below zero to access earlier elements in the array.


Growing slices (the copy and append functions)

To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.

This technique is how dynamic array implementations from other languages work behind the scenes. The next example doubles the capacity of s by making

a new slice, t, copying the contents of s into t, and then assigning the slice value t to s:

```
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x

t := make([]byte, len(s), (cap(s) + 1) * 2) // + 1 in case cap(2) == 0

for i := range s {
	t[i] = s[i]
}

s = t

```

The looping piece of this common operation is made easier by the built-in copy function. 

As the name suggests, copy copies data from a source slice to a destination slice.

It returns the number of elements copied.

> func copy(dst, src []T) int

The copy function supports copying between slices of different lengths (it will copy up to the smaller number of elements).

In addition, copy can handle source and destination slices that share the same underlying array, handling overlapping slices correctly.


```
	t := make([]byte, len(s), (cap(s) + 1) * 2)
	
	copy(t, s)
	
	s = t

```

A common operation is to append data to the end of a slice. 

This function appends byte elements to a slice of bytes, growing the slice if necessary, and returns the updated slice value:

```

	func AppendByte(slice []byte, data ...byte) []byte {
		m := len(slice)
		n := m + len(data)
		
		if n > cap(slice) { // if necessary, reallocate
			// allocate double what's needed, for future growth.
			newSlice := make([]byte, (n+1)*2)
			copy(newSlice, slice)
			slice = newSlice
		}
		slice = slice[0:n]
		copy(slice[m:n], data)
		return slice
	}

```

One could use AppendByte like this:

```
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	// p === []byte{2, 3, 5, 7, 11, 13}

```

Functions like AppendByte are useful because they offer complete control over the way the slice is grown.

Depending on the characteristics of program, it may be desirable to allocate in smaller or larger chunks, 
or to put a ceiling on the size of reallocation.

But most programs don't need complete control, so Go provides a built-in append function that's good for most
purpose; it has the signature

> func append(s []T, x ...T) []T

The append function appends the elements x to the end of the slice s, and grows the slice if a greater capacity is needed:

```
	a := make([]int, 1)
	// a == []int{0}
	a = append(a, 1, 2, 3)
	// a == []int{0, 1, 2, 3}

```

To append one slice to another, use ... to expand the second argument to a list of arguments.

```
	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

```

Since the zero value of a slice (nil) acts like a zero-length slice, you can declare a slice variable and then append to it in a loop:

```
	// Filter returns a new slice holding only
	// the elements of s that satisfy fn()
	func Filter(s []int, fn func(int) bool) []int {
		var p []int // == nil
		for _, v := range s {
			if fn(v) {
				p = append(p, v)
			}
		}
		return p
	}

```

A possible "gotcha"

As mentioned ealier, re-slicing a slice doesn't make a copy of the underlying array. The full array will be kept in 
memory until it is no longer referenced. Occasionally this can cause the program to hold all the ddata in memory when
only a small piece of it is need.

For example, this FindDigits function loads a file into memory and searches it for the first group of consecutive numeric digits, returning them as a new slice.


```
	var digitRegexp = regexp.MustCompile("[0-9]+")
	
	func FindDigits(filename string) []byte {
		b, _ := ioutil.ReadFile(filename)
		return digitRegexp.Find(b)
	}

```


This code behaves as advertised, but the returned []byte points into an array containing the entire file.
Since the slice references the original array, as long as the slice is kept around the garbage collector can't release the array;
the few useful bytes of the file keep the entire contents in memory.


To fix this problem one can copy the interesting data to a new slice before returning it:

```
	func CopyDigits(filename string) []byte {
		b, _ := ioutil.ReadFile(filename)
		b = digitRegexp.Find(b)
		c := make([]byte, len(b))
		copy(c, b)
		return c
	}

```

A more concise version of this function could be constructed by using append.

This is left as an exercise for the reader.


 