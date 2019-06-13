package main

import (
	"fmt"
	"go_dev/balance/module"
	"hash/crc32"
	"math/rand"
)

type Hash struct {
}

func init() {
	module.RegistBalance("hash", &Hash{})
}

func (this *Hash) Dobalance(insts []*module.Instance) (inst *module.Instance, err error) {
	fmt.Println("use hash balance...")
	hashKey := fmt.Sprintf("%d", rand.Int())
	table := crc32.MakeTable(crc32.IEEE)
	hashCode := crc32.Checksum([]byte(hashKey), table)
	index := int(hashCode) % len(insts)
	// fmt.Println("hashKey=", hashKey, ",table=", table, ",hashCode=", hashCode, ",index=", index)
	inst = insts[index]
	return
}
