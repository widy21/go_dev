package main

import (
	"fmt"
	"time"
)

//chanel 简单操作
func range_op() {
	var inChan chan int
	inChan = make(chan int, 10)
	for i := 0; i < 10; i++ {
		inChan <- i
	}
	close(inChan)
	for i := range inChan {
		fmt.Println(i)
	}
}

func writeChan(intChan chan int) {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
}

func readChan(intChan chan int) {
	for {
		var b, ok = <-intChan
		if !ok {
			fmt.Println("chanel cloeed...")
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println(b)
	}
}
func main() {
	//普通读写
	//range_op()
	intChan := make(chan int, 100)
	go writeChan(intChan)
	go readChan(intChan)
	//没有通信机制，通过slepp保证readChan执行
	time.Sleep(12 * time.Second)
}
