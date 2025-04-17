package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
    greeter()

	//func greeterTwo() {
	//	fmt.Println("")
//	} // not allowed to write a function inside a function

    result := adder(3, 5)
	fmt.Println("Result is:" ,  result)


	proResult, myMessage := proAdder(2,5,8,9) 
	fmt.Println("pro Result is :", proResult)
	fmt.Println("pro message is  :", myMessage)


}

func adder (valOne int , valTwo int) int  {

	return valOne + valTwo

}

func proAdder(values ...int) (int, string) {
    total := 0
	for _, val := range values{
		total += val
	}
  return total , "hii Pro Result function "
}

func greeter () {
	fmt.Println("Namaste from golang ")

}

func greeterTwo() {
    fmt.Println("Namaste from golang 2 ")

} 

