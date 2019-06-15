package main

import (
	"fmt"
)

func writeChan(inchan chan int) {
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Second)
		inchan <- i
	}
	close(inchan)
}

func readChan(inchan, outchan chan int) {
	for {
		var b, ok = <-inchan
		if !ok {
			fmt.Println("chanel cloeed...")
			break
		}
		//time.Sleep(1 * time.Second)
		//fmt.Println(b)
		outchan <- b
	}
	close(outchan)

	/*for v := range inchan {
		outchan <- v
	}
	close(outchan)*/
}
func main() {
	inChan := make(chan int, 10)
	outChan := make(chan int, 10)
	//exitChan := make(chan bool, 10)
	go writeChan(inChan)
	go readChan(inChan, outChan)
	//输出结果chanel中的值
	//go func() {
	/*for i := range outChan {
		fmt.Println("===", i)
	}*/
	for {
		v, ok := <-outChan
		if !ok {
			fmt.Println("chanel close.")
			break
		}
		fmt.Println(v)
	}
	//}()
	//time.Sleep(2 * time.Second)

}
