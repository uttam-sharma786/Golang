package main

import "fmt"

func main() {
	language := make(map[string]string)

	// To add value

	language["JS"] = "Javascirpt"
	language["Ruby"] = "Ruby"
	language["Py"] = "Python"

	fmt.Println("List of all language :", language)

	// fetch particlur element

	fmt.Println("Js short is ", language["JS"])

	// delete

	// delete(language, "RB")
	// fmt.Println("List of all language", language)

	//loop

	for key, value := range language {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
