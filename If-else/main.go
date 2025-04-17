package main

import "fmt"

func main() {
	fmt.Println("if-else in golang")

	loginCount := 23 
     var result string 


	if loginCount < 10 {
		result = "Regular user"
	} else{
		result = "something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is oddd")
	} else {
		fmt.Println("number is even")

	
}

if num := 3 ; num<10 {
   fmt.Println("num is less that 10")

} else {
	fmt.Println("num is more that 10")
}
}