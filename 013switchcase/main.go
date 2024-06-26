package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println("Switch and case in golang")

	// r := rand.New(rand.NewSource(SEED))
	// fmt.Println(r.Uint64())
	// fmt.Println(r.Uint64())

	// r := rand.New(rand.NewSource(SEED))

	// r.time.Now().UnixNano()
	// fmt.Println(r.Uint64())
	// fmt.Println(r.Uint64())

	// rand.Int(6)

	// rand.NewSource(Seed(time.Now().UnixNano()))
	diceNumber := rand.Shuffle(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again")
	default:
		fmt.Println("Wrong choice")
	}
}
