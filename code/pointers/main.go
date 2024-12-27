package main

import "fmt"

// By value
func changeNumByValue(num int) {
	num = 24
	fmt.Println("Inside changeNum:", num)
}

// By reference
func changeNumByReference(num *int) {
	*num = 24
	fmt.Println("Inside changeNum:", *num)
}

func main() {
	num := 1

	changeNumByValue(num)

	fmt.Println("After changeNumByValue call:", num)

	changeNumByReference(&num)

	// fmt.Println("Memory address:", &num)
	fmt.Println("After changeNumByReference call:", num)
}
