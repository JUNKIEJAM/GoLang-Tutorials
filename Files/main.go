package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files in GoLang")

	content := "This needs to go in the following file - sampleFile.txt"

	file, err := os.Create("./sampleFile.txt")

	if err != nil {
		panic(err)
	}

	// as per the documentation
	length, err := io.WriteString(file, content)

	fmt.Println("Length is : ", length)

	// recommended way
	defer file.Close()

	readFile("./sampleFile.txt")
}

func readFile(filename string) {

	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
}
