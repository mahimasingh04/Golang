package main 

import "fmt"

import "net/http"
import "io"

func PerformPostJSONRequest() {
    performFormRequests()
}

func PostFormRequests() {
      
	url := "http://localhost:8000/postForm"

	data := url.Values{};
	data.Add ("firstName" , "Mahima")
	data.Add("last Name", "Singh")
	data.Add("email", "mahimasingh1814@gmail.com")

	response , err := http.postForm(url, data)
    if(err !=nil) {
		panic(err)
	}
    defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Response is:", string(content))
}