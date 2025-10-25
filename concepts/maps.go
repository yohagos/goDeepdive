package concepts

import (
	"fmt"
	"reflect"
)

func Maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map => ", m)

	v1 := m["k1"]

	fmt.Println("v1 => ", v1)

	v3 := m["k3"]
	fmt.Println("v3 => ", v3)

	fmt.Println("len => ", len(m))

	/* 	delete(m, "k2")
	   	fmt.Println("map => ", m) */

	/* for k := range m {
		delete(m, k)
	}
	fmt.Println("map => ", m) */

	_, prs := m["k2"]
	fmt.Println("prs => ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map => ", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if reflect.DeepEqual(n, n2) {
		fmt.Println("n & n2 are equal")
	}
}
