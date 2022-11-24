package main

import (
	"fmt"
	"log"
	"time"
)

// Resolving the blocking operation constraint of channel
// - One solution is to receive in separate go routine
// - Another solution is to create a buffered channel by specifying capacity (in this case, 2)
func main() {
	start := time.Now()
	c := make(chan string, 2) // This allows you to fill-up a buffered channel without a corresponding receiver, and it won't block until the channel is full.
	c <- "hello"
	c <- "world"
	// c <- "three" // If you attempt to send a third value, it'll result in a deadlock since the channel buffer is set to 2 and already full.
	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
