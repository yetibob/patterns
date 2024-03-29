package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

// Message type stuff
type Message struct {
	str  string
	wait chan bool
}

// Boring it is
func Boring(msg string) chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

// FanInMessage type
func FanInMessage(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
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

// UseBoring ones
func UseBoring() {
	x := Boring("guy")
	y := Boring("girl")
	c := FanInMessage(x, y)
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}
