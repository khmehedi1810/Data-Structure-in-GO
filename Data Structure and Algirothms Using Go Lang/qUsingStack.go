package main

import (
	"fmt"
)

// ================================ Not Complete this code =============================== //
// ================================          Sorry         =============================== //

const S = 5

var s1 = make([]int, 5)
var s2 = make([]int, 5)
var top1 = -1
var top2 = -1
var count = 0

func main() {
	fmt.Println("Implement Queue using Stack: ")

	fmt.Println("1. Add Element is Queue.")
	fmt.Println("2. Delete Element from Queue. ")
	fmt.Println("Choose a option: ")
	var option int
	fmt.Scan(&option)
	//fmt.Println(option)

	switch option {
	case 1:
		fmt.Println("Enter a vale to add in stack: ")
		var num int
		fmt.Scan(&num)
		enq(&num, &count)
	case 2:
		dequeue(&top1, &top2, &count)
	default:
		fmt.Println("Invalid input. Try again.")
		main()

	}
}

func enq(num *int, count *int) {
	//fmt.Println(*num)
	push1(&*num)
	*count++
	//fmt.Println("count ", *count)
}

func push1(num *int) {

	if top1 == S-1 {
		fmt.Println("The stack is overflow.")
	} else {
		fmt.Println(*num, " is added in the stack 1")
		top1++
		s1[top1] = *num
	}
	fmt.Println("The stack 1 is: ", s1)
	main()

}
func push1a(num *int) {

	if top1 == S-1 {
		fmt.Println("The stack is overflow.")
	} else {
		fmt.Println(*num, " is added in the stack 1")
		top1++
		s1[top1] = *num
	}
	//fmt.Println("The stack 1 is: ", s1)

}
func disfinalque() {
	for i := 0; i < top1-1; i++ {
		fmt.Printf("%d ", s1[i])
	}
}

func dequeue(top1 *int, top2 *int, count *int) {
	var a1 int
	//fmt.Println("dequeue", *top1)
	if *top1 == -1 && *top2 == -1 {
		fmt.Println("The quque is Empty")
	} else {
		//fmt.Println("hi")
		for i := 0; i <= *top1+2; i++ {
			a1 = *pop1()
			push2(&a1, &*top2)
			//fmt.Println(a1)
		}
	}
	*top1 = -1
	disstack2(&*top2)

}

func pop1() *int {
	top1--
	return &s1[top1]
}

func push2(num1 *int, top2 *int) {
	if *top2 == S-1 {

	} else {
		fmt.Println(*num1, " is added in stack 2")
		*top2++
		s2[*top2] = *num1
	}
	fmt.Println("The stack 2 is: ", s2)

}
func disstack2(top2 *int) {
	for i := 0; i < *top2+2; i++ {
		push1a(&s2[i])
	}
	disfinalque()
}

//push1(&s2[i])
