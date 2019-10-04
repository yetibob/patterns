package patterns

import (
	"fmt"
	"math/rand"
)

// QuittableGen stuff
func QuittableGen(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				// perform cleanup work if necessary
				// such as delete temp files
				// cleanup()
				// quit <- True
				return
			}
		}
	}()
	return c
}

// QuitFunc stuff
func QuitFunc() {
	quit := make(chan bool)
	c := Generator("hi")
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	// if we are allowing go statements ot do some code cleanup and need to wait for them to finish
	// wait on a channel that they will use to signal they are done with work
	// Otherwise when main exits, the goroutine will end (possibly prematurely)
	// <-quit
}
