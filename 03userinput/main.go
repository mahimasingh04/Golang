package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok syntax xomes in || err err

	// there is no try catch in go lang
	input, _ := reader.ReadString('\n')  // input is try and err is catch here 
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("type of rating is %T", input) // gives string
}
