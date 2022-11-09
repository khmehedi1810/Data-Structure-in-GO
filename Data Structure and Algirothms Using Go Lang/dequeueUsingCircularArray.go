package main

import "fmt"

var lan = 5

var dequ = make([]int, lan)
var f = -1
var r = -1

func main() {
	fmt.Println("")
	fmt.Println("This is Dequeue using Circular Queue.")
	fmt.Println("1. Add data in front")
	fmt.Println("2. Add data in rear")
	fmt.Println("3. Delete Data from front")
	fmt.Println("4. Delete Data from rear")
	fmt.Println("5. Get front data")
	fmt.Println("6. Get rear data")
	fmt.Println("7. Display Data.")
	fmt.Println("Choose an option:")
	var ch int
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Println("Enter a value:")
		var o int
		fmt.Scan(&o)
		enqinFront(&o)
		//fmt.Println(o)
	case 2:
		fmt.Println("Enter a value:")
		var o int
		fmt.Scan(&o)
		enqinRear(&o)
	case 3:
		delFront()
	case 4:
		delRear()
	case 5:
		getFront()
	case 6:
		getRear()
	case 7:
		disdeq()
	default:
		fmt.Println("Invalid input. Try again.")
		main()
	}

}

func enqinFront(data *int) {

	if (f == 0 && r == lan-1) || (f == r+1) {
		fmt.Println("The Dequeue is full.")
	} else if f == -1 && r == -1 {
		f = 0
		r = 0
		fmt.Println(*data, " is added in queue.")
		dequ[f] = *data
	} else if f == 0 {
		f = lan - 1
		fmt.Println(*data, " is added in queue.")
		dequ[f] = *data
	} else {
		f--
		fmt.Println(*data, " is added in queue.")
		dequ[f] = *data
	}
	disdeq()
}

func enqinRear(data *int) {
	if (f == 0 && r == lan-1) || (f == r+1) {
		fmt.Println("The Dequeue is full.")
	} else if f == -1 && r == -1 {
		f = 0
		r = 0
		fmt.Println(*data, " is added in queue.")
		dequ[r] = *data
	} else if r == lan-1 {
		r = 0
		fmt.Println(*data, " is added in queue.")
		dequ[r] = *data
	} else {
		r++
		fmt.Println(*data, " is added in queue.")
		dequ[r] = *data
	}
	disdeq()
}

func getFront() {
	if f == -1 && r == -1 {
		fmt.Println("The Dequeue is Empty. ")
	} else {
		fmt.Println("The front value is: ", dequ[f])
	}
	main()
}

func getRear() {
	if f == -1 && r == -1 {
		fmt.Println("The Dequeue is Empty. ")
	} else {
		fmt.Println("The rear value is: ", dequ[r])
	}
	main()
}

func delFront() {
	if f == -1 && r == -1 {
		fmt.Println("The Dequeue is Empty")
	} else if f == r {
		fmt.Println(dequ[f], " is front value that deleted from the queue")
		f = -1
		r = -1
	} else if f == lan-1 {
		fmt.Println(dequ[f], " is front value that deleted from the queue")
		f = 0
	} else {
		fmt.Println(dequ[f], " is front value that deleted from the queue")
		f++
	}
	disdeq()
}

func delRear() {
	if f == -1 && r == -1 {
		fmt.Println("The Dequeue is Empty")
	} else if f == r {
		fmt.Println(dequ[r], " is rear value that deleted from the queue")
		f = -1
		r = -1
	} else if r == 0 {
		fmt.Println(dequ[r], " is rear value that deleted from the queue")
		r = lan - 1
	} else {
		fmt.Println(dequ[r], " is rear value that deleted from the queue")
		r--
	}
	disdeq()
}

func disdeq() {
	var i = f
	if f == -1 && r == -1 {
		fmt.Println("The Dequeue is Empty.")
	} else {
		fmt.Print("The Dequeue is: ")
		for i != r {
			fmt.Printf("%d ", dequ[i])
			i = (i + 1) % lan
		}
		fmt.Printf("%d ", dequ[r])
		main()
	}
}
