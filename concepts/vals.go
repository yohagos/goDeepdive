package concepts

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func Vals() {
	a, b := vals()
	fmt.Println("a, b => ", a, b)
	fmt.Println("a => ", a)
	fmt.Println("b => ", b)

	_, c := vals()
	fmt.Println("c => ", c)
}