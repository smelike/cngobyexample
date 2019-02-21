package main


import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	
	arg := os.Args[1:]
	for _, url := range arg {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("http err:")
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		//fmt.Println(os.Stdout)
	}
}
