package main

import (
	"fmt"
	"math/rand"
)

type student struct {
	name string
	age  int
	next *student
}

func transfer(p *student) {
	// 链表遍历
	for p != nil {
		fmt.Println(*p)
		// time.Sleep(time.Second)
		p = p.next
	}
}

func insertTail(p *student) {
	// 尾部插入
	tail := p
	for i := 0; i < 9; i++ {
		stu := student{
			name: fmt.Sprintf("stu%d", i),
			age:  rand.Intn(100),
		}
		tail.next = &stu
		tail = &stu
	}
}

func insertHead(p **student) {

	for i := 0; i < 9; i++ {
		stu := student{
			name: fmt.Sprintf("stu%d", i),
			age:  rand.Intn(100),
		}
		stu.next = *p
		*p = &stu
	}
}

func deleteNode(p *student) {
	// 删除指定节点
	preNode := p
	for p != nil {
		if p.name == "stu6" {
			preNode.next = p.next
			break
		}
		preNode = p
		p = p.next
	}
}

func insertNode(p *student) {
	// 在指定节点后插入新节点
	for p != nil {
		if p.name == "stu7" {
			newNode := student{"newNode", 10, nil}
			newNode.next = p.next
			p.next = &newNode
			break
		}
		p = p.next
	}
}

func main() {
	node := student{"node-stu", 1, nil}
	// insertTail(&node)

	head := &node
	insertHead(&head)
	transfer(head)
	fmt.Println("=======================")
	deleteNode(head)
	transfer(head)
	fmt.Println("=======================")
	insertNode(head)
	transfer(head)
}
