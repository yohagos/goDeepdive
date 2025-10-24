package concepts

import "fmt"

func Arrays() {
	var a [5]int
	fmt.Println("A content => ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("set: ", len(a))

	fmt.Println()
	b := [5]int{1, 20, 3, 40, 5}
	fmt.Println("dcl: ", b)

	fmt.Println()
	b = [...]int{10, 2, 30, 4, 50}
	fmt.Println("dcl: ", b)

	fmt.Println()
	b = [...]int{60, 3: 400, 500}
	fmt.Println("idx: ", b)

	var twoD [2][3]int
	for i := range make([]int, 2) {
		for j := range make([]int, 3) {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
