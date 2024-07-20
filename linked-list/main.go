package main

import "fmt"

type listNode struct {
	data int
	next *listNode
}

func main() {
	fmt.Println("invoking the driver")
	head := &listNode{data: 11}
	second := &listNode{data: 22}
	third := &listNode{data: 33}
	fourth := &listNode{data: 44}
	fifth := &listNode{data: 55}
	head.next = second
	second.next = third
	third.next = fourth
	fourth.next = fifth
	fifth.next = nil

}

func display(head *listNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%v", cur.data)
		if cur.next != nil {
			fmt.Print("->")
		}
		cur = cur.next
	}
}

// ll ops
// insert(x)
// delete(x)
// lookups(x)
// list()/Display()
