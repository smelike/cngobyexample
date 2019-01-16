package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)
	
	const n = 50000000
	
	const d = 3e20 / n // 3e20 = 3*10^20
	fmt.Println(d)
	
	fmt.Println(int64(d))
	
	fmt.Println(math.Sin(n))
}