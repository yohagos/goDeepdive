package concepts

import "fmt"

func Loops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range make([]int, 5) {
		fmt.Println("range ", i)
	}
	fmt.Println()
	for {
		fmt.Println("loop and break")
		break
	}
	fmt.Println()
	for n := range make([]int, 6) {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Modulo by 2 | at index ", n)
	}
}
