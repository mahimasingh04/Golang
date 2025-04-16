package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[2] = "watermelon"
	fruitList[3] = "mango"
fmt.Println("Fruit list is: ", fruitList)
fmt.Println("Fruit list is: ", len(fruitList)) //strange

var vegList = [3]string{"potato",  "tomato", "beans"}
fmt.Println("Vegy list is:", len(vegList))

}