package module

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegistBalance("radom", &RandomBalance{})
}

func (this *RandomBalance) Dobalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no insts")
		return
	}
	instsNum := len(insts)
	index := rand.Intn(instsNum)
	inst = insts[index]
	return
}
