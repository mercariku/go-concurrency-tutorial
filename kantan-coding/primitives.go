package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	c := make(chan string)
	d := make(chan string)
	go someFunc("1") // buffered channel is asynchronous (this means the sending go routine can send and forget) - meanwhile, unbuffered channel is synchronous because sending and receiving take place in synch.
	go func() {
		c <- "test" // put string data onto the channel, c
	}()

	go func() {
		d <- "data"
	}()

	time.Sleep(time.Second * 2)
	msg := <-c
	fmt.Println(msg)
	fmt.Println("hi")

	select {
	case msgOne := <-c:
		fmt.Println(msgOne)
	case msgTwo := <-d:
		fmt.Println(msgTwo)
	}

}

// unbuffered channel is synchronous because the sending side is blocked until there's a response from the receiver
// buffered channel is asynchronous because sending go routine can send data to the buffer queue (until it's full). The sending go routine can keep processing workload without waiting for response by a receiver
// the sending go routine is blocked once the queue is full
