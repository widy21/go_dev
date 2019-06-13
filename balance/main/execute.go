package main

import (
	"flag"
	"fmt"
	"go_dev/balance/module"
	"math/rand"
	"time"
)

func main() {
	var hosts []*module.Instance
	for i := 0; i < 16; i++ {
		tempInst := module.CreateInstance(fmt.Sprintf("127.0.%d.%d", rand.Intn(255), rand.Intn(255)), 8080)
		hosts = append(hosts, tempInst)
	}
	fmt.Println(hosts)

	var balancerName string
	flag.StringVar(&balancerName, "b", "", "no balncer name after [-b]")
	flag.Parse()
	fmt.Printf("[-b] balancerName=%s \n", balancerName)
	fmt.Printf("=================\n")
	for {
		inst, error := module.Dobalance(balancerName, hosts)
		if error != nil {
			fmt.Println("error: ", error)
		}
		time.Sleep(time.Second)
		fmt.Println("inst == ", inst)
	}
}
