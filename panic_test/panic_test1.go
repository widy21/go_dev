package main

//测试panic

import "fmt"

func badCall() {
	panic("bad call...")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("error occur: %v\n", e)
		}
	}()
	fmt.Println("before call...")
	badCall()
	fmt.Println("after call...")
}

func main() {
	test()
}
