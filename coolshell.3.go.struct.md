> 结构体

Go 的结构体和 c 的基本一样，不过在初始化时有些不一样，Go 支持带名字的初始化。

```
type Person struct {
  name string
  age int
  email string
}

func main() {
  // 初始化
  person := Person{"Tom", 30, "tom@gmail.com"}
  person = Person{name:"Tom", age: 30, email: "jack@gmail.com"}
  
  fmt.Println(person)
  
  pPerson := &person // 传引用？
  fmt.Println(pPerson) // {Tom 30 jack@gmail.com}
  
  pPerson.age = 46
  person.name = "Julie"
  fmt.Println(person) // 输出 {Julie 46 jack@gmail.com}
}
```

----

> 结构体方法

注意：Go 语言中没有 `public, protected, private` 的关键字，所以，**如果你想让一个方法可以被别的包访问的话，
你需要把这个方法的【第一个字母大写】**。*这是一种约定。*

```
type rect struct {
  width, height int
}

// 定义结构体的 area 方法
func (r *rect) area() int { // 求面积
  return r.width * r.height
}

// 定义结构体的 perimeter 方法
func (r *rect) perimeter() int {
  return 2 *(r.width + r.height)
}

func main() {
  // 给结构体 rect 赋值
  r := rect(r.width + r.height)
  fmt.Println("Area:", r.area())
  fmt.Println("Perimeter:", r.perimeter())
  
  rp := &r
  fmt.Println("Area:", rp.area())
  fmt.Println("Perimeter:", rp.perimeter())
}

```
