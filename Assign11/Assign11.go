/* Homework - 10
JSON serialization and deserialization using Go
References: 
http://blog.golang.org/json-and-go
*/

package main

import (
    "encoding/json"
	"fmt"
)

type Friend struct{
	Name string
	Description string
	Place string
}

func main() {

	// Encoding
    f1 := Friend{"Ritu", "Best Pizza-Maker", "Confused Indian Guy"}
	b1, err := json.Marshal(f1)
	if(err == nil){
	fmt.Println("The encoded JSON object is "+string(b1))
	}
	
	// Decoding with incomplete JSON object
	b2 := []byte(`{"Name":"Radhika","Place":"Pune"}`) // Description field missing
	var f2 Friend
	json.Unmarshal(b2, &f2)
	fmt.Println("The decoded incomplete JSON object is Name: "+f2.Name+" Description: "+f2.Description+" Place: "+f2.Place)
	
	// Decoding with complete JSON object
	b3 := []byte(`{"Name":"Sunil","Description":"Yellow Shirt Guy","Place":"Bangalore"}`) // All fields present
	var f3 Friend
	json.Unmarshal(b3, &f3)
	fmt.Println("The decoded complete JSON object is Name: "+f3.Name+" Description: "+f3.Description+" Place: "+f3.Place)
}
