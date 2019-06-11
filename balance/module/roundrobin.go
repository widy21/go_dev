package module

import (
	"errors"
	"math/rand"
)

type RoundRobin struct {
}

func (this *RoundRobin) Dobalance(insts []*Instance) (inst *Instance, err error) {
	//todo ...
	if len(insts) == 0 {
		err = errors.New("no insts")
		return
	}
	instsNum := len(insts)
	index := rand.Intn(instsNum)
	inst = insts[index]
	return
}
