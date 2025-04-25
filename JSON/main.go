package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // this is aliasing
	Price    int
	Platform string `json:"website"`
	Password string
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json")
	EncodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs bootcamp", 299, "lco", "abc123", []string{"web dev", "js"}},
		{"advanced backend ", 499, "lco", "bcd123", []string{"full stack", "nodejs"}},
		{"blockchain", 599, "lco", "abce2123", []string{"web 3", "solidity"}},
		{"Angular bootcamp", 299, "lco", "iobc123", []string{"javascript", "frontend framework"}},
	}

	// package this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
