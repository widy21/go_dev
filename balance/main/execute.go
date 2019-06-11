package main

import (
	"fmt"
	"go_dev/balance/module"
	"math/rand"
)

func main() {
	var hosts []*module.Instance
	for i := 0; i < 16; i++ {
		tempInst := module.CreateInstance(fmt.Sprintf("127.0.%d.%d", rand.Intn(255), rand.Intn(255)), 8080)
		hosts = append(hosts, tempInst)
	}
	fmt.Println(hosts)
}
