package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Tomato"
	fruits[3] = "Peach"

	fmt.Println("Fruit list is : ", fruits)
	fmt.Println("Size of fruits is ", len(fruits))

	var vegList = [5]string{"potato", "beans", "musroom"}
	fmt.Println("Veg list is: ", len(vegList))

}
