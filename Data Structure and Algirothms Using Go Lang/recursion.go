package main

import "fmt"

func main() {
	var res = fact(5)
	fmt.Println("Result is: ", res)
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + fact(n-1)
	}
}
