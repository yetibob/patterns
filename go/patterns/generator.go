package patterns

import "fmt"

// Generator docs
// A Generator is a function that creates a channel, kicks off a go routine that will send/recieve
// on that channel, and then returns the channel to you
// This is useful for creating "services"
// Where for example you create run multiple instances of the generator
// to act upon multiple differing sets of inputs/outputs
// Example Usage
// channel := generator("hi guy")
// for { fmt.Println(<-channel) }
func Generator(msg string) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return c
}
