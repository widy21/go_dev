package main

import (
	"fmt"
	"sync"
	"time"
)

/**
goroute 计算阶乘。
*/

type task struct {
	n int
}

var (
	m    = make(map[int]uint64)
	lock sync.Mutex
)

func cal(n int) {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	lock.Lock()
	m[n] = uint64(sum)
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		tempTask := &task{
			n: i,
		}
		go cal(tempTask.n)
	}

	//output result...
	time.Sleep(time.Second * 10)

	//cal(4)

	lock.Lock()
	for k, v := range m {
		fmt.Printf("num[%d]'s result is[%d]\n", k, v)
	}
	lock.Unlock()
}
