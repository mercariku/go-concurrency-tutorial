package main

import (
	"fmt"
)

func main() {
	go count("sheep")
	count("fish")
}
func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
	}
}
