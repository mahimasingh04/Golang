
package main 

import "fmt"
func main() {
	defer fmt.Println("how are you ?")
	defer fmt.Println("Raunit");
	fmt.Println("Hello");
	myDefer();
}


func myDefer () {
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
}
//last in first out