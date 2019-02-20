// Unix 里 echo 命令的一份实现，echo 把它的命令行参数打印成一行。 
package main


import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
