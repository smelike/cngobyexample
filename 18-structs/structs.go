package main

import "fmt"

type person struct {
	name string
	age int
}


func main() {
	
	fmt.Println(person{"Bob", 20})
	
	fmt.Println(person{name: "Alice-wonder", age: 30})
	
	fmt.Println(person{name: "Fred"})
	
	fmt.Println(&person{name: "Anna", age: 40})
	
	s := person{name: "Sean-Gina", age: 50}
	fmt.Println(s.name)
	
	sp := &s
	fmt.Println(sp.age)
	
	sp.age = 56
	fmt.Println(sp.age)
	fmt.Println(sp)
}