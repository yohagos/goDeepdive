package concepts

import "fmt"

func sum(nums ...int) {
	fmt.Println("received nums ")
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total sum => ", total)
}

func Variadic() {
	sum(1, 2, 3)
	sum(5, 17, 18, 30)

	nums := []int{2, 4, 6, 8, 10}
	sum(nums...)
}
