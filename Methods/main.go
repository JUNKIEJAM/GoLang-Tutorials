package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")

	prithesh := User{
		"Prithesh",
		"abc@gmail.com",
		true,
		15}

	prithesh.GetStatus()
}

func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)

}
