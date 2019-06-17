package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum - 1)
	go test()
	for i := 0; i < 2; i++ {
		go cal()
	}
	time.Sleep(100 * time.Second)
}

func cal() {
	for {
		fmt.Println("is calling...")
		time.Sleep(time.Second)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test error...", err)
		}
	}()

	var m map[string]int
	m["test"] = 1
}
