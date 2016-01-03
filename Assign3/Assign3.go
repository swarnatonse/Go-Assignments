/* Homework - 3
I have written a simple Go program that uses go-routines and channels
The program takes three strings and concurrently reverses them and puts the result in a channel
In the main function, the three output strings are read from the channel and displayed

References:
http://tour.golang.org/concurrency/2
http://rosettacode.org/wiki/Reverse_a_string#Go
*/

package main
 
import (
    "fmt"
)
 
func reverseBytes(s string, ch chan string) {
    r := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        r[i] = s[len(s)-1-i]
    }
    ch <- string(r)
}

func main() {
	ch := make(chan string)
    go reverseBytes("anrawS", ch)
	go reverseBytes("iruag", ch)
	go reverseBytes("Swarnagauri", ch)
	
	str1, str2, str3 := <-ch, <-ch, <-ch
	fmt.Println("The reversed strings are:")
	fmt.Println(str1, str2, str3)
}

/* Output:
The reversed strings are:
iruaganrawS Swarna gauri
*/