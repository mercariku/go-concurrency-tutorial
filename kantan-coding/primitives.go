package main

import (
	"fmt"
	"log"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	start := time.Now()

	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	elapsed := time.Since(start)
	log.Printf("%s", elapsed)

}
