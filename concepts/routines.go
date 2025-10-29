package concepts

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.

func f3(from string) {
	for i := range 3 {
		fmt.Println(from, " => ", i)
	}
}

func Routines() {
	// Suppose we have a function call f(s). Here’s how we’d
	// call that in the usual way, running it synchronously.
	f3("direct")

	// To invoke this function in a goroutine, use go f(s).
	// This new goroutine will execute concurrently with the calling one.
	go f3("go routine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate
	// goroutines now. Wait for them to finish (for a more robust
	// approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")
}
