package main

import (
	"fmt"
	"sort"
)

func main() {
	i := 100
	var arr = [...]int{1, 2, 3, 5, 01}
	sort.Ints(arr[:])
	fmt.Println("hello world~", i, arr)
}
