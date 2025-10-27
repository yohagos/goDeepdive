package concepts

import "fmt"

// Go’s structs are typed collections of fields. 
// They’re useful for grouping data together to form records.

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name.
func newPerson(name string) *person {
	// Go is a garbage collected language; you can safely return a 
	// pointer to a local variable - it will only be cleaned up by 
	// the garbage collector when there are no active references to it.
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 40})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 24})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 34})

	// It’s idiomatic to encapsulate new struct creation in 
	// constructor functions
	fmt.Println("new Person => ", newPerson("John"))

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 35}
	fmt.Println("S name => ", s.name)

	// You can also use dots with struct pointers - the pointers 
	// are automatically dereferenced.
	sp := &s
	fmt.Println("sp age => ", sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println("sp age => ", sp.age)

	// If a struct type is only used for a single value, 
	// we don’t have to give it a name. The value can have an 
	// anonymous struct type. This technique is commonly used 
	// for table-driven tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Tim",
		true,
	}
	fmt.Println("dog => ", dog)
}
