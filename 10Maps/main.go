package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["CPP"] = "C Plus plus"

	fmt.Println("list of all langauges: ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages", languages)

   //loops are interesting in golangs
   for key, value := range languages {
	fmt.Println("For key %v, value is %v", key, value)
	
   }
   
}
