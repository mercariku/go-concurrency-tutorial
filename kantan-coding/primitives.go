package main

import (
	"fmt"
	"log"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func doWork(done <-chan bool) { // this is how you declare read only channel
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work")
		}
	}
}

func main() {
	start := time.Now()

	// Done channel pattern
	// for select pattern is common because of the done channel
	done := make(chan bool)
	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)

	elapsed := time.Since(start)
	log.Printf("%s", elapsed)

}
