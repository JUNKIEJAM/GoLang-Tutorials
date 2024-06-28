package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// fmt.Println("")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	content, _ := io.ReadAll((response.Body))
	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(string(content))

}
