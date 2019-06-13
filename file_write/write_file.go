package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	file, error := os.OpenFile("/tmp/gotest", os.O_RDWR|os.O_CREATE, 755)
	if error != nil {
		fmt.Println("op file error: ", error)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(fmt.Sprintf("%d\n", rand.Int()))
	}
	writer.Flush()
}
