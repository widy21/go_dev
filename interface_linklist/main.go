package main

import "fmt"

func main() {
	var link Link
	for i := 0; i < 10; i++ {
		// link.InsertHead(i)
		link.InsertTail(fmt.Sprintf("node-%d", i))
	}
	link.trans()
}
