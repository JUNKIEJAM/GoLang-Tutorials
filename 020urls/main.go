package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params are: %T\n", qparams)

	for _, val := range qparams {
		fmt.Println("Para is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh"}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
