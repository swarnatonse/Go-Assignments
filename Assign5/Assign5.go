/* Homework - 5
Implementing Producer-Consumer problem in a round-about way using mutexes from "sync" package and 
an implementation of semaphores
I'm so sorry I could not get this program to work by the deadline, both goroutines enter a deadlock and I couldn't figure out why :'(
I am also attaching another program of the Producer-Consumer problem using channels without mutexes or semaphores
References:
https://jlmedina123.wordpress.com/2014/04/08/255/
https://golang.org/pkg/sync/
http://www.golangpatterns.info/concurrency/semaphores
https://golang.org/pkg/math/rand/#example__rand
http://blog.golang.org/go-slices-usage-and-internals
*/

package main
import (
    "fmt"
    "math/rand"
    "sync"
)

type semaphore chan int
var buffer []int
var buffer_empty semaphore
var buffer_full semaphore
var mutex = &sync.Mutex{}

func P (s semaphore, n int){
	var e = 0
    for i := 0; i < n; i++ {
        s <- e
    }
}

func V (s semaphore, n int) { 
    for i := 0; i < n; i++ {
        <-s
    }
}

func AddToBuffer(value int){
	if len(buffer) >= 4{
		fmt.Println("Buffer full!")
		}else{
		buffer = append(buffer, value)
		fmt.Println("Added to buffer!")
		}
}

func RemoveFromBuffer() int{
	if len(buffer) == 0{
		fmt.Println("Buffer empty!")
		return -1 }else{
		value:=buffer[0]
		buffer = append(buffer[:0], buffer[1:]...)
		return value}
}

func producer(){
	fmt.Println("Inside producer")
	r := rand.New(rand.NewSource(99))
	for i:=0; i<10; i++{
	fmt.Println("Entered here")
	value := r.Intn(10)
	V(buffer_full,1)
	mutex.Lock()
	fmt.Println("Critical section")
	AddToBuffer(value)
	mutex.Unlock()
	P(buffer_empty,1)
	fmt.Println("Producer added %d to the buffer!", value)
		
	}
}

func consumer(){
	fmt.Println("Inside consumer")
	for i:=0; i<10; i++{
	V(buffer_empty,1)
	mutex.Lock()
	value:=RemoveFromBuffer()
	mutex.Unlock()
	P(buffer_full,1)
	fmt.Println("Consumer removed %d from buffer!", value)
	}
}

func main(){
	fmt.Println("Inside main")
	buffer = make([]int, 4)
	buffer_empty = make(semaphore, 1)
	buffer_full = make(semaphore, 1)

	go producer()
	go consumer()
}