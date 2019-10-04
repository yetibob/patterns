package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

// Result type for use with dummy google search
type Result struct {
	msg string
}

var (
	// Web Search
	Web = FakeSearch("web")
	// Image Search
	Image = FakeSearch("image")
	// Video Search
	Video = FakeSearch("video")
)

// Search type
type Search func(query string) Result

// FakeSearch sweetness
func FakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

// Google search thing
func Google(query string) (results []Result) {
	// this is slow because it requries orderly completion of each request
	// results = append(results, Web(query))
	// results = append(results, Image(query))
	// results = append(results, Video(query))

	// do with fan in pattern to execute queries at the same time
	c := make(chan Result)
	go func() { c <- Replicate(query, Web, 3) }()
	go func() { c <- Replicate(query, Image, 3) }()
	go func() { c <- Replicate(query, Video, 3) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// Replicate will perform searches via "service" replication
func Replicate(query string, replica Search, count int) Result {
	c := make(chan Result)
	for i := 0; i < count; i++ {
		go func() {
			c <- replica(query)
		}()
	}
	return <-c
}

// DoSearch stuff
func DoSearch() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}
