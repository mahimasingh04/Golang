package main

import "fmt"

func main() {
	fmt.Println(("Welcome to a class onn pointers"))

	//var ptr *int

	//fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual ponter is", ptr)
	fmt.Println("Value of actual ponter is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is : "  , myNumber)

}
