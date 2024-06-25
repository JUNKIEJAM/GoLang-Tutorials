package mai

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	prithesh := User {
		"Prithesh",
		"abc@gmail.com",
		true,
		15
	}
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
