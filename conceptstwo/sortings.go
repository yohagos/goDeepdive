package conceptstwo

import (
	"fmt"
	"slices"
)

// Go’s slices package implements sorting for builtins and
// user-defined types. We’ll look at sorting for builtins first.

func Sortings() {
	strs := []string{"c", "a", "b"}
	// Sorting functions are generic, and work for any ordered
	// built-in type. For a list of ordered types, see cmp.Ordered.
	slices.Sort(strs)
	fmt.Println("Strings => ", strs)

	ints := []int{7, 2, 4}
	// An example of sorting ints.
	slices.Sort(ints)
	fmt.Println("Ints => ", ints)

	// We can also use the slices package to check if a slice is
	// already in sorted order.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted => ", s)
}
