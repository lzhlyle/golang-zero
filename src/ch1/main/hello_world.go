package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	//fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1:])
	}

	//os.Exit(-1)
}
