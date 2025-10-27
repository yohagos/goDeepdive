package concepts

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.

// Here’s a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we’ll implement this interface on rect and circle types.
type rects struct {
	width, height float64
}

type circles struct {
	radius float64
}

// To implement an interface in Go, we just need to implement
// all the methods in the interface. Here we implement geometry
// on rects.
func (r rects) area() float64 {
	return r.width * r.height
}

func (r rects) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.
func (c circles) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circles) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here’s a generic
// measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Sometimes it’s useful to know the runtime type of an interface
// value. One option is using a type assertion as shown here; another
// is a type switch.
func detectCircle(g geometry) {
	if c, ok := g.(circles); ok {
		fmt.Println("circle with radius => ", c.radius)
	}
}

func Interfaces() {
	r := rects{width: 3.5, height: 6.4}
	c := circles{radius: 5.6}

	// The circle and rect struct types both implement the geometry
	// interface so we can use instances of these structs as arguments
	// to measure.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
