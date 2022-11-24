package main

import "fmt"

// Demonstrating a constraint of channel - Blocking operation
func main() {
	c := make(chan string)
	c <- "hello" // 1: Here, we send a value through the channel, c, which blocks the next operation.

	msg := <-c // 2: since the previous operation (send block) is blocking this operation (receive block), msg is unable to receive resulting in a deadlock.
	fmt.Println(msg)
}
