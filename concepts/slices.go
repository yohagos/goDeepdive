package concepts

import (
	"fmt"
)

func Slices() {
	var s []string
	fmt.Println("unint => ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp => ", s, " len => ", len(s), " cap => ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set => ", s)
	fmt.Println("get => ", s[2])

	fmt.Println("len => ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd => ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy => ", c)

	l := s[2:5]
	fmt.Println("sl1 => ", l)

	l = s[:5]
	fmt.Println("sl2 => ", l)

	l = s[2:]
	fmt.Println("sl3 => ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl => ", t)

	t2 := []string{"g", "h", "i"}
	/* if slices.Compare(t, t2, ) == 0 {
		fmt.Println()
	} */
	fmt.Println("dcl => ", t2)

	twoD := make([][]int, 3)
	for i := range make([]int, 3) {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range make([]int, innerLen) {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
