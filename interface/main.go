package main

import "fmt"

type Car interface {
	Run()
	GetName() string
	Whistle()
}

type Test interface {
	Hello()
}

type BYD struct {
	Name string
}

func (p *BYD) Run() {
	fmt.Println("byd is run.")
}
func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Whistle() {
	fmt.Println("I'm whistle on five hundred miles away.")
}

func (p *BYD) Hello() {
	fmt.Println("in hello()...")
}

func main() {
	var a Car
	// var b int
	// a = b
	// fmt.Printf("type a is %T\n", a)
	var byd BYD
	byd.Name = "byd1"
	a = &byd
	fmt.Println("the name of the car is ", a.GetName())
	a.Run()
	a.Whistle()

	var test Test
	test = &byd
	test.Hello()
}
