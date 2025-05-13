package main

import ( 
	"fmt"
	 "encoding/json"
	 
)

type course struct {
	Name     string  `json:"coursename"` // this is aliasing
	Price    int
	Platform string  `json:"website"`
	Password string 
	Tags     []string  `json:"tags,omitempty"`
}
func main() {
   fmt.Println("welcome to json")
   DecodeJSON()
}
func EncodeJson() {
    lcoCourses := []course{
        {"Reactjs bootcamp", 299, "lco", "abc123", []string{"web dev", "js"}},
        {"advanced backend ", 499, "lco", "bcd123", []string{"full stack", "nodejs"}},
        {"blockchain", 599, "lco", "abce2123", []string{"web 3", "solidity"}},
        {"Angular bootcamp", 299, "lco", "iobc123", []string{"javascript","frontend framework"}} , 
    }

    // package this data as json data
    finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", finalJson)
}

func DecodeJSON(){
   jsonDataFromWeb := []byte(`
   {

    "coursename" : "ReactJS Bootcamp",
	"Price" : "299",
	"website" : "lco",
	"tags": ["web dev","js"]
	  }
   `)

   var lcoCourse course
   checkValid := json.Valid(jsonDataFromWeb)

   if checkValid {
	fmt.Println("JSON was valid")
	json.Unmarshal(jsonDataFromWeb, &lcoCourse)
	fmt.Println("%#v\n", lcoCourse)
   } else {
	fmt.Println("JSON was not valid")
   }

   var myOnlineData map[string]interface{}
   json.Unmarshal(jsonDataFromWeb, &myOnlineData)
   fmt.Println("%#v\n", myOnlineData)

   for k, v := range myOnlineData {
	fmt.Println("key is %v an value is %v and Type is %T", k, v, v)

   }
}