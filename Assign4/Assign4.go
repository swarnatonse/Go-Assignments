/* Homework - 4
Experimenting with built-in hash table functionality of Go called "map"
This is a simple program that makes use of maps 
This is a program to find the most compatible "friends" among a list of people
Friendship is decided based on a language the person speaks. 
If a person on the list speaks the language you desire, the person is listed as a friend 
References:
https://blog.golang.org/go-maps-in-action
http://blog.golang.org/go-slices-usage-and-internals
https://golang.org/doc/effective_go.html
http://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line */

package main
 
import (
   "fmt"
)

type Person struct{
	Name string
	Language []string
}

func findFriends(speaks map[string][]Person, searchStr string) []string{
// This function takes the input language as argument an searches the map for friends
// Unfortunately search is case sensitive - possible improvement
	var listfriends []string
	for _, p := range speaks[searchStr] {
        listfriends = append(listfriends, p.Name)
    }
	return listfriends
}

func main(){

var friends [4]Person // Should have tried using pointer []*Person - possible improvement

// Initializing list of 4 potential friends and the languages they speak
// Could have used better initializing technique - possible improvement
friends[0].Name = "Arvindh"
friends[0].Language = []string{"English", "Hindi", "Tamil", "Kannada"}

friends[1].Name = "Vikaasa"
friends[1].Language = []string{"English", "Tamil", "Telugu"}

friends[2].Name = "Amritesh"
friends[2].Language = []string{"English", "Hindi", "Telugu"}

friends[3].Name = "Sunil"
friends[3].Language = []string{"English", "Hindi", "Kannada", "Marathi"}

// This is to populate the map with the list of people that speak a language
// The language string is the Key for this Map
speaks := make(map[string][]Person)
    for _, p := range friends {
        for _, l := range p.Language {
            speaks[l] = append(speaks[l], p)
        }
	}


var input string
fmt.Println("Enter a language that you speak")
fmt.Scanln(&input) // This input does not take spaces

foundfriends := findFriends(speaks, input)

// Display friends. No names are displayed if no friends are found
fmt.Println("Your friends are:")
for _, i:= range foundfriends{
	fmt.Println(i)}
}