/* Homework - 12
I know this is highly irresponsible of me but I almost forgot about this assignment and I have an AoA exam in 2 hours, 
So here is a simple Go Program experimenting with time package and ticks 
This will hopefully not be the last of my experiments with Go 
References:
https://tour.golang.org/concurrency/6
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	rockYou := time.Tick(1000 * time.Millisecond)
	message := time.After(4000 * time.Millisecond)
	for {
		select {
		case <-rockYou:
			fmt.Println("ROCK YOU!")
		case <-message:
			fmt.Println("<3 <3 <3 DO MORE GO, GO IS LOVE, GO IS WAY BETTER THAN CLASS SUMMARIES, DOS IS THE BEST <3 <3 <3")
			return
		default:
			fmt.Println("We will...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
