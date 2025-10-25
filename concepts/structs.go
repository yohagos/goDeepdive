package concepts

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	fmt.Println(person{"Bob", 40})
	fmt.Println(person{name: "Alice", age: 24})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 34})

	fmt.Println("new Person => ", newPerson("John"))

	s := person{name: "Sean", age: 35}
	fmt.Println("S name => ", s.name)

	sp := &s
	fmt.Println("sp age => ", sp.age)

	sp.age = 51
	fmt.Println("sp age => ", sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Tim",
		true,
	}
	fmt.Println("dog => ", dog)
}
