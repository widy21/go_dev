package main

import "fmt"

type student struct {
	name  string
	age   int
	left  *student
	right *student
}

func transfer(p *student) {
	// 二叉树遍历
	if p == nil {
		return
	}
	fmt.Println(*p)
	transfer(p.left)
	transfer(p.right)
}

func main() {
	root := student{"root", 1, nil, nil}
	left1 := student{"left1-stu", 1, nil, nil}
	left2 := student{"left2-stu", 1, nil, nil}
	right1 := student{"right1-stu", 1, nil, nil}
	root.left = &left1
	left1.left = &left2
	root.right = &right1
	transfer(&root)
}
