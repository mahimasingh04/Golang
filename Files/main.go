
package main 

import ("fmt"
"os"
"io")


func main () {
	fmt.Println("welcome to files in golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	os.Create("./mygofile.txt")
	file, err := os.OpenFile("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)
	defer file.Close()

}

