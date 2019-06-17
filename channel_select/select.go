package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 chan int
	var ch2 chan int
	ch1 = make(chan int, 10)
	ch2 = make(chan int, 10)
	go func() {
		var i int
		for {
			ch1 <- i
			time.Sleep(time.Second)
			ch2 <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()

	/*for i := 0; i < 10; i++ {
		ch1 <- i
		ch2 <- i * i
	}*/

	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("get data time out...")
			return

			/*default:
			fmt.Println("get data time out...")
			time.Sleep(time.Second)*/
		}

	}

}
