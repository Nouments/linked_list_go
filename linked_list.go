package linkedlist

import "fmt"

type linkedList struct {
	info string
	next *linkedList
}

// Create a new linked list
func createLinkedList(info string) *linkedList {
	return &linkedList{
		info: info,
		next: nil,
	}
}

// inserting value into the right of the linked list
func (c *linkedList) insertRight(info string) *linkedList {
	p := createLinkedList(info)
	for c.next!=nil{
		c = c.next
	}
	c.next=p
	return c
}

//inserting value into the left of the linked list

func (c *linkedList) insertLeft(info string) *linkedList {
	n:=createLinkedList(info)
	n.next=c
	return n
}

// show from left->right the value
func (c *linkedList) showLinkedList() {
	if c.next != nil {
		fmt.Printf("%v->",c.info)
		c.next.showLinkedList()
	}else{
		fmt.Printf("%v",c.info)
	}
	
}
