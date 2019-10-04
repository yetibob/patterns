package main

import (
	"fmt"
	"sync"
)

func main() {
	x := []int{1, 2, 3, 4}

	var wg sync.WaitGroup
	wg.Add(len(x))

	for _, num := range x {
		go func(num int) {
			defer wg.Done()
			fmt.Printf(" I have recieved x: %v\n", num)
		}(num)
	}
	wg.Wait()
}
