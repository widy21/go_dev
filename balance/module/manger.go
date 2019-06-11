package module

import (
	"errors"
	"fmt"
)

type manager struct {
	allBalancer map[string]Balancer
}

var mgr = &manager{
	allBalancer: make(map[string]Balancer),
}

//私有方法
func (this *manager) registBalance(name string, item Balancer) {
	this.allBalancer[name] = item
}

//公共方法
func RegistBalance(name string, item Balancer) {
	mgr.registBalance(name, item)
}

//公共方法
// func DoBalance(name string, insts []*Instance) (ins *Instance, err error) {
func Dobalance(name string, insts []*Instance) (ins *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = errors.New(fmt.Sprintf("not find balance: %s", name))
		return
	}
	ins, err = balancer.Dobalance(insts)
	return
}
