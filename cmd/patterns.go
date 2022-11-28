package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
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

type Spinlock struct{
	state *int32
}

const free = int32(0)

func(1 *Spinlock) Lock(){
	for !atomic.CompareAndSwapInt32(1.state, free, 42) {// 42 or any other value but 0
		runtime.Gosched()  // Poke the scheduler
	}
}

func(1 *Spinlock) Unlock(){
	atomic.StoreInt32(1.state, free) // Once atomic, always atomic!
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

//5// Where channels fail
- You can create deadlocks with channels
- Channels pass around copies, which can affect performance
- Passing pointers to channels can create race conditions
- What about "naturally shared" structures like caches or registries? (DON'T do this!!)

//5.1// Mutexes
- One potential solution to problems caused by channel is 'Mutexes'
- But...
- Mutexes are like toilets
- The longer you occupy them, the longer the queue gets
- Read/Write mutexes can only reduce the problem
- Using multiple mutexes will cause deadlocks sooner or later
- All-in-all, not the solution we are looking for

//5.2// Three Shades of Code
- Blocking = Your program may get locked up (for undefined time)
- Lock free = At least one part of your program is always making progress
- Wait free = All parts of your program are always making progress
- How to write lock free or wait free code?

//5.3// Atomic operations
- sync.atomic package
- Store, Load, Add, Swap and CompareAndSwap
- Mapped to thread-safe CPU instructions
- These instructions only work on integer types
- Only about 10-60x slower than their non-atomi counter-parts

//5.4// Spinning CAS
- You need a 'state' variable and a 'free' constant
- Use CAS (CompareAndSwap) in a loop:
-- If state is not free, try again until it is
-- If state is free, set it to something else
- If you managed to change the state, you 'own' it

//6// Ticket storage
- We need an indexed data structure, a 'ticket' and a 'done' variable
- A fucntion draws a new ticket by adding 1 to the ticket
- Every ticket number is unique as we never decrement
- Treat the ticket as an index to store your data
- Increase done to extend the 'ready to read' range

*/
