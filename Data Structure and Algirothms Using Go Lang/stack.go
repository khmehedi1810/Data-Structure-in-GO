package main

import "fmt"

const M = 5

var top = -1
var sta [M]int

func main() {
	fmt.Println("Welcome to the stack world: ")
	var n int

	fmt.Println("1. Add a value in stack. ")
	fmt.Println("2. Delete a value. ")
	fmt.Println("3. Display the Stack. ")
	fmt.Println("Choose your option: ")
	fmt.Scan(&n)

	switch n {
	case 1:
		var x int
		fmt.Println("Enter a value: ")
		fmt.Scan(&x)
		push(&x, &top)
		//fmt.Println(x)
	case 2:
		pop(&top)
	case 3:
		//fmt.Println("This ", n)
		dis(&top)
	default:
		fmt.Println("Choose an right option.")
		main()
	}
}

func push(x *int, top *int) {

	if *top == M-1 {
		fmt.Println("Overflow")
	} else {
		*top++
		sta[*top] = *x
		fmt.Println(*x, "is added in stack")
	}
	main()
}

func pop(top *int) {
	fmt.Println("Pop section: ")
	if *top == -1 {
		fmt.Println("UnderFlow")
	} else {
		fmt.Println(sta[*top], "Deleted from the stack")
		*top--
	}

	main()
}

func dis(top *int) {
	fmt.Println("Display section: ")
	if *top == -1 {
		fmt.Println("Underflow")
	} else {
		for i := *top; i >= 0; i-- {
			fmt.Printf("%d ", sta[i])
		}
		fmt.Println("")
	}

	main()
}
