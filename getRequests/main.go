package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main () {
	fmt.Println("Welcome to the server")
	PerformGetRequests()
}

func PerformGetRequests() {
	const myUrl= "http://localhost:3000/get-json"  // should use something else

	response, err := http.Get(myUrl)
	if(err != nil) {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("status lentgh is :", response.ContentLength)

      

	//byteCount, _ := reponseString.Write(content)
	// fmt.Println(string(content))

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)


	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is :",  byteCount)
	fmt.Println(responseString.String())
	


}