package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("there is no arg...")
		return
	}
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
