package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type student struct {
	Name string
	Id   string
	Age  int
}

type studentArray []student

// Len is the number of elements in the collection.
func (p studentArray) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p studentArray) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

// Swap swaps the elements with indexes i and j.
func (p studentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var stus studentArray
	for i := 0; i < 10; i++ {
		stu := student{
			Name: fmt.Sprintf("stu-%d", rand.Intn(100)),
			Id:   fmt.Sprintf("110%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Printf("\n................after sorted.............\n")
	sort.Sort(stus)

	for _, v := range stus {
		fmt.Println(v)
	}
}
