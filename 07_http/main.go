package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	// Error control breaks but ingnoring doesn't
	// i, err2 := resp.Body.Read(bs)
	// if err2 != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Read: ", i)
	fmt.Println(string(bs))
}
