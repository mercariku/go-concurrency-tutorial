package main

import (
	"fmt"
	"time"
)

/*

Refer to the "Concurrency Patterns in Go" video [Apr 16, 2018] by Arne Claus
https://www.youtube.com/watch?v=YEKjSzIwAdA&t=29

*/

func main() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c) //1// Output: 0, false

}

func TryReceive(c <-chan int) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	default: // processed when c is blocking
		return 0, true, false
	}
}

func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration): // time.After() returns a channel
		return 0, true, false
	}
}

func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In {
		select {
		// send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
}

func Turnout(InA, InB <-chan int, OutA, OutB chan int) {

	for {
		select { // receive from first non-blocking
		case data, more = <-InA:
		case data, more = <-InB:
		}
		if !more {
			return nil
		}
		select { // send to first non-blocking
		case OutA <- data:
		case OutB <- data:
		}

	}
}

func TurnoutWithQuitChannel(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	for {
		select {
		case data = <-InA:
		case data = <-InB:
		case <-Quit:
			close(InA) // remember that close generates message
			close(InB) // also remember that this is actually an anti-pattern (but you could argue that 'Quit' acts as a delegate

			Fanout(InA, OutA, OutB) // flush the remaining data
			Fanout(InA, OutA, OutB) // flush the remaining data
			return
		}
	}
}

/*

//1//
- a receive always returns two values
- 0 as it is the zero value of int
- fase because "no more data" or "returned value is no valid"

//2//
- Channels are stream of data
- Dealing with multiple streams is the true power of select
- Only close channel from sending side, never (with some exceptions) from the receiving side

//3//
- Data Stream Shape
- Fan-out
- Funnel
- Turn-out

//4//
- 90% of use cases can be covered with CSP and select

*/
