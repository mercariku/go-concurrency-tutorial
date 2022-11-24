package main

import (
	"fmt"
	"time"
)

func main() { // 1: main routing starts
	c := make(chan string)
	go count("sheep", c) // 2: second go routine starts which runs the count function

	for msg := range c { // 4: exists the loop once c range satisfied
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // 3: send value through channel created in line 9 (remember that sending through channel is a blocking operation)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
