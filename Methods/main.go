package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
     mahima := User{"mahima", "mahi@go.dev", true , 21} 
	 fmt.Println(mahima)   //  // {mahima mahi@go.dev true 21}
	 fmt.Printf("mahima details are: %+v\n", mahima) // more structure and detailed format mein lautata h - mahima details are: {Name:mahima Email:mahi@go.dev Status:true Age:21}
	fmt.Printf("name is %v and email is %v", mahima.Name, mahima.Email)
    mahima.GetStatus()
	mahima.NewMail()
	fmt.Printf("name is %v and email is %v", mahima.Name, mahima.Email)  // this is to check whether its making a copy or actiually changing the value
}

type User struct { 
	Name   string
	Email  string
	Status bool
	Age    int
}


func (u User ) GetStatus() {
   fmt.Println("Is User active", u.Status)

}

func (u User) NewMail() {
	u.Email ="test@go.in"
	fmt.Println("email of user is : ", u.Email)
} 