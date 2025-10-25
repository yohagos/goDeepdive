package concepts

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	println("initial => ", i)

	zeroval(i)
	println("zeroval => ", i)
	zeroptr(&i)
	println("zeroptr => ", i)

	println("pointer => ", &i)
}
