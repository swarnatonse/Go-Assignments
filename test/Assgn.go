/* DOS Homework-2
I have implemented a Stable Marriage Problem code in Go
Stable Marriage algorithm is a simple resource allocation algorithm 
Algorithm:
function stableMatching {
    Initialize all m ∈ M and w ∈ W to free
    while ∃ free man m who still has a woman w to propose to {
       w = highest ranked woman to whom m has not yet proposed
       if w is free
         (m, w) become engaged
       else some pair (m', w) already exists
         if w prefers m to m'
           (m, w) become engaged
           m' becomes free
         else
           (m', w) remain engaged
    }
}
This example is taken from http://www.geeksforgeeks.org/stable-marriage-problem/
Other references:
https://en.wikipedia.org/wiki/Stable_marriage_problem
http://golang.org/doc/code.html
https://tour.golang.org/
*/


package main

import (
	"fmt"
)

const(
	N = 4 	// Number of men and women
)

func wPreference(prefer [][] int, w int, m int, m1 int) bool{
// If woman prefers m1 over m, don't break current engagement
	for i := 0; i < N; i++ {
		if prefer[w][i] == m1{
			return true
		}
		
		if prefer[w][i] == m{
			return false
		}
	}
	return true
}

func stableMarriage(prefer [][] int) []int{
	var wPartner[N] int
	var mTaken[N] bool
	
	for i:=0; i < N; i++{
		wPartner[i] = -1	// initialize partners for all women to -1 
		mTaken[i] = false}	// and taken status of all men to false
		
	freeCount := N

	for freeCount > 0{
	
	var m int
	for m = 0; m < N; m++{
		if mTaken[m] == false{ 
		break }
	}
	// 
	for i:=0; i<N && mTaken[m] == false; i++{
		var w int
		w = prefer[m][i]
		// If woman is free, pair her up with man of preference
		if wPartner[w-N] == -1{
			wPartner[w-N] = m
			mTaken[m] = true
			fmt.Printf("%d is paired with %d\n", w,m)
			freeCount--
		} else{	// if woman is not free, check if current man is of higher preference and break engagement if so
		var m1 int
		m1 = wPartner[w-N]
		if wPreference(prefer, w, m, m1) == false{
			wPartner[w-N] = m
			mTaken[m] = true
			mTaken[m1] = false
			fmt.Printf("%d is paired broken up with %d and paired with %d\n", w, m1, m)
		}
		}
	}
	}
	fmt.Println("Woman	Man")	// Print the result. Maybe should have tried returning the array and printing in main
	for i:=0; i < N; i++{
		fmt.Printf("%d\t%d\n", i+N, wPartner[i])}
	return wPartner
}

func main() {
	
	prefer := make([][]int, N*2) // Making a slice with the preferences of the women
	for i := range prefer {
		prefer[i] = make([]int, N)}
	prefer = [][]int{	// List of preferences
	[]int{7, 5, 6, 4},
	[]int{5, 4, 6, 7},
	[]int{4, 5, 6, 7},
	[]int{4, 5, 6, 7},
	[]int{0, 1, 2, 3},
	[]int{0, 1, 2, 3},
	[]int{0, 1, 2, 3},
	[]int{0, 1, 2, 3}}
	pairings := stableMarriage(prefer)
	for i:= range pairings{
		fmt.Printf("%d\t%d\n", i+N, pairings[i])}
}