package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all language : ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages : ", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key %v, value is %v", value)
	}

}
