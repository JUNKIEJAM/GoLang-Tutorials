package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Banana", "Papaya"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Grapes")
	fmt.Println(fruitList)

	// slicing at index 1
	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// slicing at index 1 util index 3
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// HIGH SCORES

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 956
	highScores[2] = 231
	highScores[3] = 943

	fmt.Println(highScores)

	highScores = append(highScores, 3454, 2345, 163)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// removing a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "ruby", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	// [reactjs javascript ruby python]

	fmt.Println(courses)

}
