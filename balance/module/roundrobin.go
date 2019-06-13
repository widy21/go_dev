package module

import (
	"errors"
	"fmt"
)

type RoundRobin struct {
	currentIndex int
}

func init() {
	RegistBalance("RoundRobin", &RoundRobin{})
}

func (this *RoundRobin) Dobalance(insts []*Instance) (inst *Instance, err error) {
	fmt.Println("use RoundRobin balancer...")
	if len(insts) == 0 {
		err = errors.New("no insts")
		return
	}
	instsNum := len(insts)
	if this.currentIndex == instsNum {
		this.currentIndex = 0
	}
	index := this.currentIndex
	inst = insts[index]
	this.currentIndex++
	return
}
