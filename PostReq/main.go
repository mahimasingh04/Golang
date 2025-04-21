package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to learn how ")
}

func PerformPostRequest() {
	const myUrl = "hhtp://localhost:3000/post"

	//fake json payload
	requestBody := strings.NewReader (`
	{
        "coursename " :"Lets go with golang",
		"price": 0,
		"platform" :"100xdevs"
		
	}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)

	if(err != nil) {
		panic(err)
	}
      defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}