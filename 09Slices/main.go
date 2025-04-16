package main

import (
	"fmt"
	"sort"
)
func main() {
   var fruitList = []string{"guava", "seetaPhal"}

   fruitList = append(fruitList, "Mango", "Banana")

   fmt.Println(fruitList)

   fruitList = append(fruitList[1 : ])
   fmt.Println(fruitList)

   //fruitList = append(fruitList[1 : 3])
   //fmt.Println(fruitList)

  // fruitList = append(fruitList[:3]) 
  // fmt.Println(fruitList)

   highScores := make([]int,  4)

   highScores[0] = 254
   highScores[1] = 325
   highScores[2] = 435
   highScores[3] = 975
   //highScores[4] = 777 error dega
   fmt.Println(highScores)

   highScores = append(highScores, 555, 889, 880)  // it will not give error because all values will be accomodated
 fmt.Println(highScores)

 sort.Ints(highScores)
fmt.Println(sort.IntsAreSorted(highScores))  // boolean 
sort.Ints(highScores)
fmt.Println(highScores)

var courses = []string {"reactjs" , "nodejs", "swift", "python", "ruby"}

fmt.Println(courses)
 var index int = 2;
 courses = append(courses [:index], courses[index+1 :] ...);
 fmt.Println(courses)

} 
 