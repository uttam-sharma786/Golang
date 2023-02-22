package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	// no inheritance in golang, no super parent or child

	Uttam := User{"Uttam", "Uttam@gmail.dev", true, 19}
	fmt.Println(Uttam)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
