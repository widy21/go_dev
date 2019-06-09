package main

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (link *Link) InsertHead(node interface{}) {
	link_node := &LinkNode{
		data: node,
		next: nil,
	}
	if link.head == nil && link.tail == nil {
		link.tail = link_node
		link.head = link_node
		return
	}

	link_node.next = link.head
	link.head = link_node
}

func (link *Link) InsertTail(node interface{}) {
	link_node := &LinkNode{
		data: node,
		next: nil,
	}
	if link.head == nil && link.tail == nil {
		link.tail = link_node
		link.head = link_node
		return
	}

	link.tail.next = link_node
	link.tail = link_node
}

func (p *Link) trans() {
	current_position := p.head
	for current_position != nil {
		fmt.Println(current_position.data)
		current_position = current_position.next
	}
}
