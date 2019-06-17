package main

import (
	"fmt"
)

func writeChan(inchan chan<- int, exitChan chan bool) {
	for i := 0; i < 10; i++ {
		inchan <- i
	}
	close(inchan)
	exitChan <- true
}

func readChan(inchan <-chan int, exitChan chan bool) {
	for {
		var b, ok = <-inchan
		if !ok {
			fmt.Println("channel close")
			break
		}
		fmt.Println(b)
	}
	exitChan <- true
}
func main() {
	inChan := make(chan int, 10)
	exitChan := make(chan bool, 10)
	go writeChan(inChan, exitChan)
	go readChan(inChan, exitChan)

	//等待协程执行完成。
	for i := 0; i < 2; i++ {
		<-exitChan
		fmt.Println("waite goroute ", i, " exited")
	}

}
