package patterns

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

// DaisyChain to test in main
func DaisyChain() {
	const n = 50
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 0 }(right)
	fmt.Println(<-leftmost)
}
