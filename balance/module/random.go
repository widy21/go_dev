package module

import (
	"errors"
	"fmt"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegistBalance("random", &RandomBalance{})
}

func (this *RandomBalance) Dobalance(insts []*Instance) (inst *Instance, err error) {
	fmt.Println("use random balancer...")
	if len(insts) == 0 {
		err = errors.New("no insts")
		return
	}
	instsNum := len(insts)
	index := rand.Intn(instsNum)
	inst = insts[index]
	return
}
