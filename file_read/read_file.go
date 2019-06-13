package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(path string) {
	file, error := os.Open(path)
	if error != nil {
		fmt.Println("open file error: ", error)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, error := reader.ReadString('\n')
		if error == io.EOF {
			fmt.Println("file end...")
			break
		}
		if error != nil {
			fmt.Println("read line error: ", error)
			break
		}
		fmt.Println(str)
	}

}
func main() {
	readFile("/tmp/test2.txt")
}
