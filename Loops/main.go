package main

import "fmt"

func main() {
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for index, day := range days {
		fmt.Println("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 5 {
			goto tutorial
			// break
		}

		fmt.Println("Value is : ", rogueValue)
		rogueValue++
	}

tutorial:
	fmt.Println("Jump to this line.")

}
