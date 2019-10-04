package patterns

import (
	"fmt"
	"time"
)

// FanIn Do it
// This pattern builds upon the generator pattern previously discussed
// The fan in pattern takes multiple inputs and pipes their outputs into a single channel
// This allows the channel receiver to act on the channels in a not blocking/out of order way
func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// UseFanIn stuff
func UseFanIn() {
	x := Generator("hi")
	y := Generator("bye")
	piped := FanIn(x, y)
	for i := 0; i < 10; i++ {
		fmt.Println(<-piped)
	}
}

// FanInSelect is the same thing but using a single goroutine plus select to avoid waiting
func FanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// UseFanInSelect Example Usage
func UseFanInSelect() {
	x := Generator("hi")
	y := Generator("bye")
	piped := FanInSelect(x, y)
	// Or if we wanted to act on a total timeout
	// timeout := time.After(5 * time.Second)
	// then we would select on timeout directly
	for i := 0; i < 10; i++ {
		select {
		case s := <-piped:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Too slow")
		}
	}
}
