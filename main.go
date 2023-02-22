package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"
const myurl string = "https://courses.learncodeonline.in/learn/Complete-Javascript-course"

func main() {
	fmt.Println("LCO wbe requests")

	fmt.Println(myurl)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
