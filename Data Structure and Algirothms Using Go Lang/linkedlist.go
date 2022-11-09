package main

import "fmt"

type node struct {
	val  int
	next *node
}
type list struct {
	start *node
	prev  *node
	temp  *node
}

func main() {

	s1 := list{}
	fmt.Println("")
	fmt.Println("Do you wanna insert a value?")
	var get int
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Scan(&get)
	if get == 1 {
		s1.insert()
		dislist(&s1)
	} else {
		fmt.Println("Good luck")
	}
}

func (l *list) insert() {
	fmt.Println("Insert Data")
	var data int
	fmt.Scan(&data)
	if l.start == nil {
		l.start.val = data
		l.start.next = nil
		l.prev = l.start
	} else {
		l.temp.val = data
		l.prev = l.start
		for l.prev.next != nil {
			l.prev = l.prev.next
		}
		l.temp.next = nil
		l.prev.next = l.temp
	}
	return
}

func dislist(s1 *list) {
	fmt.Println("List is: ")
	head := s1.start
	if head == nil {
		fmt.Println("Empty list")
	} else {
		for head.next != nil {
			fmt.Printf("%d ", head.val)
			head = head.next
		}
	}
}
