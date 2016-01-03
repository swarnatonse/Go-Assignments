/* Homework - 8
I have written a Go program that runs a function that is kind of like the requirement from Project 1 and which uses SHA-256 encoding.
From the requirement pdf-
"Output: Print, on independent lines, the input string and the corresponding
SHA256 hash separated by a TAB, for each of the bitcoins you find. Obviously,
your SHA256 hash must have the required number of leading 0s (k = 3 means
3 0's in the hash notation). An extra requirement, to ensure every group finds
different coins, it to have the input string prefixed by the gatorlink ID of one of
the team members."
k is hardcoded to 1 
References:
https://en.wikipedia.org/wiki/SHA-2
https://golang.org/pkg/crypto/sha256/
https://golang.org/pkg/bytes/
http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
https://tour.golang.org/concurrency/1
*/

package main

import (
	"crypto/sha256"
	"fmt"
	"bytes"
	"math/rand"
	"time"
)

const(
	k = 1
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Function to generate random strings
func RandStringBytes(n int) []byte {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return b
}

// Function that does the SHA-256 encoding and checks the leading zeroes
func KindOfLikeTheBitCoinThing(substring []byte, actorID int){
	for i := 0; i < 1000; i++{
	randomString := RandStringBytes(8)
	gatorID := "swarnatonse;"
	var finalString []byte
	finalString = append([]byte(gatorID),randomString...)
	hashedValArr := sha256.Sum256(finalString)
	hashedVal := hashedValArr[:]
	if bytes.HasPrefix(hashedVal, substring) {
	fmt.Printf("Thread %d came up with %s %x\n", actorID, finalString, hashedVal)}
	}
	}
	
func main() {
	
	substring := make([]byte, k)
	for i:= range substring{
	substring[i] = 0}
	fmt.Println(substring)
	go KindOfLikeTheBitCoinThing(substring, 0)
	go KindOfLikeTheBitCoinThing(substring, 1)
	go KindOfLikeTheBitCoinThing(substring, 2)
	go KindOfLikeTheBitCoinThing(substring, 3)
	time.Sleep(10 * time.Second) // Adding time for go routines to run
}
