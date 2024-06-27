package main

import "fmt"

// nested fuctions are not allowed

func main() {
	fmt.Println("Welcome to functions in golang")
	val, str := func1(2, 3, 4, 6)

	fmt.Println("The sum is : ", val)
	fmt.Println(str)
}

func func1(arr ...int) (int, string) {

	total := 0

	for _, value := range arr {

		total += value
	}

	return total, "Sum done"
}
