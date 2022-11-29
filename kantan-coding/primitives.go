package main

import (
	"fmt"
	"log"
	"time"
)

func sliceToChannel(nums []int) <-chan int { // the "<-chan int" just means "it returns a read-only channel"
	out := make(chan int) // notice it's an unbuffered channel
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

var pl = fmt.Println
var lp = log.Printf

func main() {
	start := time.Now()

	// Pipeline

	nums := []int{2, 3, 4, 7, 1} // input

	// stage 1
	dataChannel := sliceToChannel(nums) // here, we put each element of 'nums' onto a channel (i.e. dataChannel)
	// stage 2
	finalChannel := sq(dataChannel)

	// Final stage
	for n := range finalChannel {
		pl(n)
	}
	elapsed := time.Since(start)
	lp("%s", elapsed)

}
