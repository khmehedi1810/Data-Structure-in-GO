package main

import (
	"fmt"
)

var arr [10]int
var enSize int

func main() {
	fmt.Println("")
	fmt.Println("Welcome to array.")
	fmt.Println("1. Add Elements in array")
	fmt.Println("2. Slow Array Elements")
	fmt.Println("3. Insertion")
	fmt.Println("4. Deletion")
	fmt.Println("5. Searching")

	var chose int
	fmt.Println("Enter your option: ")
	fmt.Scan(&chose)

	switch chose {
	case 1:
		addArr()
	case 2:
		disArr(&enSize)
	case 3:
		var choos int
		fmt.Println("A. First Insertion")
		fmt.Println("B. Middle Insertion")
		fmt.Println("C. Last Insertion")
		fmt.Println("Choose your Insertion. ")
		fmt.Scan(&choos)
		switch choos {
		case 1:
			var datum int
			fmt.Println("Enter the data you want to insert.")
			fmt.Scan(&datum)
			firInsert(&datum)
		case 2:
			var datum int
			fmt.Println("Enter the data you want to insert.")
			fmt.Scan(&datum)
			midInsert(&datum)
			//disArr(&dat)
		case 3:
			var datum int
			fmt.Println("Enter the data you want to insert.")
			fmt.Scan(&datum)
			lastInsert(&datum)
		default:
			fmt.Println("Invalid Input. Try again")
			main()
		}
	case 4:
		var delData int
		fmt.Println("Enter the position you want to delete. ")
		fmt.Scan(&delData)
		delitem(&delData)
	case 5:
		fmt.Println("Enter the data you want to search. ")
		var xxxx int
		fmt.Scan(&xxxx)
		search(&xxxx, &enSize)
	default:
		fmt.Println("Invalid Input. Try again")
		main()
	}

}

func search(data *int, enSixe *int) {
	//var check = 0
	for i := 0; i < *enSixe; i++ {
		//fmt.Printf("%d ", arr[i])
		if *data == arr[i] {
			fmt.Println("The data is found at position ", i+1)
			main()
		}
	}

	fmt.Println("Data is not found")

	main()
}

func delitem(delpos *int) {
	if *delpos <= 0 && *delpos > enSize {
		fmt.Println("Invalid. Try again")
		main()
	} else {
		for i := *delpos - 1; i < enSize; i++ {
			arr[i] = arr[i+1]
		}
		enSize--
	}
	disArr(&enSize)

}

func addArr() {
	fmt.Println("Enter the array size.")
	fmt.Scan(&enSize)
	fmt.Println("Enter elements for Array: ")

	for i := 0; i < enSize; i++ {
		fmt.Scan(&arr[i])
	}

	disArr(&enSize)
}

func midInsert(data *int) {
	fmt.Println("Enter the position you want to insert.")
	var pos int
	fmt.Scan(&pos)
	if pos <= 0 && pos >= len(arr) {
		fmt.Println("Invalid Position. Try again")
		main()
	} else {
		for i := enSize; i >= pos; i-- {
			arr[i+1] = arr[i]
		}
		arr[pos] = *data
	}
	enSize = enSize + 1
	disArr(&enSize)
}

func firInsert(data *int) {
	if enSize == 10 {
		fmt.Println("Array is full.")
		main()
	} else if arr[0] == 0 {
		arr[0] = *data
	} else {
		for i := enSize; i >= 0; i-- {
			arr[i+1] = arr[i]
		}
		arr[0] = *data
	}
	enSize = enSize + 1
	disArr(&enSize)
}

func lastInsert(data *int) {
	if enSize == 10 {
		fmt.Println("Array is full.")
		main()
	} else {
		arr[enSize] = *data
	}
	enSize = enSize + 1
	disArr(&enSize)
}

func disArr(sizeA *int) {
	for i := 0; i < *sizeA; i++ {
		fmt.Printf("%d ", arr[i])
	}
	main()
}
