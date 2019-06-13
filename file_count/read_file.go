package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Countx struct {
	CharCount  int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func countFile(path string) {
	var cx Countx
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
		//统计字符个数
		runeArray := []rune(str)
		for _, v := range runeArray {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				cx.CharCount++
			case v == ' ' || v == '\t':
				cx.SpaceCount++
			default:
				cx.OtherCount++
			}
		}
	}
	fmt.Println("file char count = ", cx.CharCount)
	fmt.Println("file num count = ", cx.CharCount)
	fmt.Println("file space count = ", cx.SpaceCount)
	fmt.Println("file other count = ", cx.OtherCount)

}
func main() {
	countFile("/tmp/test2.txt")
}
