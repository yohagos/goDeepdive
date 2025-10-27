package concepts

import "fmt"

func Loops() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	
	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic “do this N times” 
	// iteration is range over an integer.
	// there was an issue with using a regulare int
	for i := range make([]int, 5) {
		fmt.Println("range ", i)
	}
	
	// for without a condition will loop repeatedly until you 
	// break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop and break")
		break
	}
	
	// You can also continue to the next iteration of the loop.
	for n := range make([]int, 6) {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Modulo by 2 | at index ", n)
	}
}
