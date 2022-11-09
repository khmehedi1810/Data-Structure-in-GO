package main

import "fmt"

const N = 5

var font = -1
var rear = -1
var que [N]int

func main() {
	fmt.Println("Queue: ")
	var n int
	fmt.Println("Choose your option: ")
	fmt.Println("1. Add a value in queue. ")
	fmt.Println("2. Delete a value. ")
	fmt.Println("3. Display the Queue. ")
	fmt.Scan(&n)

	switch n {
	case 1:
		var a int
		fmt.Println("Enter a value: ")
		fmt.Scan(&a)
		//fmt.Println("The value is: ", a)
		enque(&a, &font)
	case 2:
		deq(&font)
	case 3:
		display(&font)
	default:
		fmt.Println("Not Valid")
	}
}

func enque(x *int, font *int) {
	//fmt.Println("The value is:", *x)
	if rear >= N-1 || *font == N-2 {
		fmt.Println("Overflow")
	} else if *font == -1 && rear == -1 {
		*font++
		rear++
		que[rear] = *x
		fmt.Println(*x, " value is added.")
	} else {
		rear++
		que[rear] = *x
		fmt.Println(*x, " value is added.")
	}

	//display
	main()

}
func deq(font *int) {

	if *font == -1 && rear == -1 {
		fmt.Println("Underflow")
	} else {
		fmt.Println(que[*font], "is deleted.")
		*font++
		main()
	}

	//display
	// for i := *font; i < N-1; i++ {
	// 	fmt.Printf("%d ", que[i])
	// }
	// main()
	//display(font)
}

func display(font *int) {
	if *font == -1 && rear == -1 {
		fmt.Println("Underflow")
	} else {
		for i := *font; i <= rear; i++ {
			fmt.Printf("%d ", que[i])
		}
	}
	main()
}
