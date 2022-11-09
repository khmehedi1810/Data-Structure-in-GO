package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

var len = 0
var head *Node

func NewNode(value int) *Node {
	var n Node
	n.value = value
	n.next = nil
	return &n
}
func TraverseLinkedList(head *Node) {
	fmt.Printf("Linked List: ")
	temp := head
	for temp != nil {
		fmt.Printf("%d -> ", temp.value)
		temp = temp.next
	}
	fmt.Printf("NULL")
	main()
}
func main() {

	fmt.Println("Enter data: ")
	var datt int
	fmt.Scan(&datt)
	//head := NewNode(datt)
	if len == 0 {
		head = NewNode(datt)
		len++
	} else {
		tempo := head
		for tempo.next != nil {
			tempo = tempo.next
		}
		tempo = NewNode(datt)
	}
	TraverseLinkedList(head)
}
