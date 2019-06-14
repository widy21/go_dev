package main

import (
	"fmt"
	"sort"
)

func main() {
	i := 100
	j := 100
	var arr = [...]int{1, 2, 3, 5, 0}
	sort.Ints(arr[:])
	fmt.Println("arr = ", arr)
	fmt.Println("hello world~", i, arr)
	fmt.Println(j)
}
