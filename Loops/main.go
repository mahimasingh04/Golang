package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)
    for d := 0; d< len(days); d++{
	fmt.Println(days[d])
    

	for i := range days{
		fmt.Println(days[i])
	}

	for _,  day := range days{
		fmt.Println("index is %v and value is %v\n", day)

	}

	rougueValue := 1
 for rougueValue < 10 {


	if rougueValue == 2 {
		goto lco 
	}
	if rougueValue == 5 {
		rougueValue ++ 
		continue 
	}
   fmt.Println("value is :", rougueValue)
   rougueValue ++
 }
}

 lco : 
   fmt.Println("Jumping at learnCodeOnline.in")
}
