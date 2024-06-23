package main

import (
	"fmt"
)

func main() {
	var username string = "prithesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallFloat float64 = 255.34565432345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicity type
	var website = "wwww.google.com"
	fmt.Println((website))

	// no var style
	numberOfUser := 300000.00
	fmt.Println(numberOfUser)

}
