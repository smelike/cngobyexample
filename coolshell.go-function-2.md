> 函数

Go 声明变量类型和函数返回值，是反过来的。

```
package main

import "fmt"

func max(a int, b int) int { // 注意参数和返回值是怎么声明的
  if a > b {
    return a
  }
  return b
}

func main() {
  fmt.Println(max(4, 6))
}
```

----

> 函数返回多个值

Go 中很多 package 都会返回两个值，一个是正常值，一个是错误，如下所示：

``
package main

import "fmt"

func main() {
  v, e := multi_ret("one")
  fmt.Println(v, e) // 输出 1 true
  
  v, e := multi_ret("four")
  fmt.Println(v, e) // 输出 0 false
  
  // 通常用法（注意分号后有 e）
  if v, e = multi_ret("four"); e {
    // 正常返回
  } else {
    // 出错返回
}

func multi_ret(a int, b int) (int, bool) {
  m := map[string]int{"one": 1, "tow": 2, "three": 3}
  var err bool
  var val int
  
  val, err = m[key]
  
  return val, err
}
```

----

> 函数不定参数

```
func sum(nums ...int) {
  fmt.Println(nums, " ") // 输出如 [1, 2, 3] 之类的数组
  total := 0
  
  for _, num = range nums { // 要的是值 不是下标
    total += num
  }
  
  fmt.Println(total)
}

func main() {
  sum(1, 2)
  sum(1, 2, 6)
  
  nums := []int{1, 2, 3, 4}
  sum(nums..)
}

```

----

> 函数闭包

 nextNum 这个函数返回了一个匿名函数，这个匿名函数记住了 nextNum 中 `i+j` 的值，并改变了 `i, j`的值，于是形成了一个闭包的用法：
 
 ```
 func nextNum() func() int {
  i, i := 1, 1
  return func int {
    var tmp = i + j
    i, j = j, tmp
    return tmp
  }
 }
 
 // main 函数中对 nextNum 的调用，主要是打出下一个斐波拉契数
 
 func main() {
  nextNumFunc := nextNum()
  for i := 0; i < 10; i++ {
    fmt.Println(nextNumFunc())
  }
 }
 ```
 
 ----
 
 > 函数的递归
 
 和 c 基本是一样的
 
 ```
 func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
 }
 
 func main() {
  fmt.Println(fact(7))
 }
 
 ```
