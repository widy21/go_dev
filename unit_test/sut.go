package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type student struct {
	Name string
	Age  int
}

func (p *student) save() (err error) {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("data marshal error:", err)
		return
	}
	err = ioutil.WriteFile("/tmp/stu.dat", data, 0755)
	return
}

func (p *student) load() (err error) {
	data, err := ioutil.ReadFile("/tmp/stu.dat")
	if err != nil {
		fmt.Println("ReadFile error:", err)
		return
	}
	err = json.Unmarshal(data, p)
	return
}
