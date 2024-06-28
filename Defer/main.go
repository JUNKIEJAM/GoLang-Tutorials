package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	fmt.Println("Two")
	fmt.Println("Three")
	helper()
}

func helper() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Output :

// Two
// Three
// 4
// 3
// 2
// 1
// 0
// One
// World
