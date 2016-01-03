/* Homework - 5 supplement
This is the Producer-Consumer problem using channels
Buffered channels are an excellent way to synchronize between two goroutines

References:
https://tour.golang.org/concurrency/3
http://www.golangpatterns.info/concurrency/producer-consumer
*/

package main

import ("fmt")

var release = make(chan bool)
var add_buffer = make(chan int)

func producer () {
	fmt.Println("Inside producer")
    for i := 0; i < 10; i++ {
		fmt.Println("Adding ",i," to buffer!")
        add_buffer <- i
    }
    release <- true
}

func consumer () {
	fmt.Println("Inside consumer")
    for {
      item := <-add_buffer
      fmt.Println("Got ",item," from buffer!")
   }
}

func main () {
   go producer()
   go consumer()
   <- release
}