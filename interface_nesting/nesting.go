package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWrite interface {
	Reader
	Writer
}

type file struct {
}

func (p *file) Read() {
	fmt.Println("reading")
}

func (p *file) Write() {
	fmt.Println("writing")
}

func Test(rw ReadWrite) {
	rw.Read()
	rw.Write()
}

func main() {
	var f file
	Test(&f)

	//测试是否属于一个接口类型
	var empty_interface interface{}
	//因为是指针类型实现的接口，所以要把指针赋值给空接口，再判断接口类型的时候才是true
	empty_interface = &f
	if v, ok := empty_interface.(ReadWrite); ok {
		fmt.Println(v, ok)
	}
	// v, ok := empty_interface.(ReadWrite)
	// fmt.Println(v, ok)
}
