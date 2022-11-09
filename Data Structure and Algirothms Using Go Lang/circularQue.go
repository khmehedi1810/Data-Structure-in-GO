package main

import (
	"fmt"
)

var cir = make([]int, 5)
var fnt = -1
var reear = -1

func main() {
	fmt.Println("This is Circular Queue: ")
	var choise int
	fmt.Println("1. Add element is queue.")
	fmt.Println("2. Delete Element from the queue. ")
	fmt.Println("3. Display. ")
	fmt.Println("Enter your choice: ")
	fmt.Scan(&choise)
	switch choise {
	case 1:
		var deta int
		fmt.Println("Enter a value")
		fmt.Scan(&deta)
		enqueue(&deta, &fnt, &reear)
	case 2:
		deque(&fnt, &reear)
	case 3:
		disCirQ(&fnt, &reear)
	default:
		fmt.Println("Invalid Input. Try again.")
		main()
	}

}

func enqueue(data *int, fnt *int, reear *int) {
	//fmt.Println("The value is: ", *data)
	if *fnt == -1 && *reear == -1 {
		fmt.Println(*data, " is added in the queue. ")
		*fnt = 0
		*reear = 0
		cir[*reear] = *data
	} else if (*reear+1)%5 == *fnt {
		fmt.Println("The queue is overflow.")
	} else {
		*reear = (*reear + 1) % 5
		cir[*reear] = *data
	}
	disCirQ(&*fnt, &*reear)
}

func deque(fnt *int, reear *int) {
	if *fnt == -1 && *reear == -1 {
		fmt.Println("The queue is Empty.")
	} else if *fnt == *reear {
		*fnt = -1
		*reear = -1
	} else {
		fmt.Println(cir[*fnt], " is deleted.")
		*fnt = (*fnt + 1) % 5
	}
	disCirQ(&*fnt, &*reear)
}

func disCirQ(fnt *int, reear *int) {
	var i = *fnt
	fmt.Print("The queue is: ")
	if *fnt == -1 && *reear == -1 {
		fmt.Println("The queue is Empty.")
	} else {
		//fmt.Println(cir)
		for i != *reear {
			fmt.Printf("%d ", cir[i])
			i = (i + 1) % 5
		}
		fmt.Printf("%d ", cir[*reear])
	}
	fmt.Println("")
	main()
}
