package main

import "fmt"
  // jwtToken := is not allowed 
  // var jwtToken = "hhjj" is allowed globally 

  const LoginToken = "hiithere" // this is public starting with capital letter (L)
func main() {
	var username string = "Mahima Singh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n ", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n ", smallValue)

	var smallFloat float32 = 255.455
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n ", smallFloat)

	//default values and some aliases 
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T\n", anotherVariable)

	// implicit styles
	var website = "ok"
	fmt.Println(website)

	// no var styles

	numberOfUser := 30000  // inside a method only not globally 
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of tpe : %T \n", LoginToken)
}
