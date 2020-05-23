> 错误处理 - Error 接口

函数错误返回可能是 C/C++ 中最让人纠结的东西，Go 的多只返回可以让我们更容易的返回错误，其可以在返回一个常规的返回值之外。
还能轻易地返回一个详细的错误描述。通常情况下，错误的类型是 error，它有一个内建的接口，

```
type error interface {
  Error() string
}
```

看示例：

```
package main

import "fmt"
import "errors"

type myError struct {
  arg int
  errMsg string
}

// 实现 Error 接口
func (e *myError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}

// 两种出错
func error_test(arg int) (int, error) {
  if arg < 0 {
    return -1, errors("Bad Arguments - negative")
  } else if arg > 256 {
    return -1, &myError{arg, "Bad Arguments - too large!"}
  }
  return arg*arg, nil
}

func main() {
  for _, i := range []int{-1, 4, 1000} {
    if r, e := error_test(i); e != nil {
      fmt.Println("failed", e)
    } else {
      fmt.Println("success:", r)
    }
  }
}
```
