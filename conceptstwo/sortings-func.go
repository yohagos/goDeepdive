package conceptstwo

import (
	"cmp"
	"fmt"
	"slices"
)

// Sometimes we’ll want to sort a collection by something other
// than its natural order. For example, suppose we wanted to sort
// strings by their length instead of alphabetically. Here’s an
// example of custom sorts in Go.

func SortingsFunc() {
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths.
	// cmp.Compare is helpful for this.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Now we can call slices.SortFunc with this custom comparison
	// function to sort fruits by name length.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// We can use the same technique to sort a slice of values
	// that aren’t built-in types.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 23},
		{name: "Alex", age: 72},
	}

	// Sort people by age using slices.SortFunc.
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)

	// HINT!! Note: if the Person struct is large, you may want the
	// slice to contain *Person instead and adjust the sorting
	// function accordingly. If in doubt, benchmark!
}
