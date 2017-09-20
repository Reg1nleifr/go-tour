package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// Hello World, 世界 = Shìjiè
	fmt.Printf("Hello, 世界\n")

	// First Package
	fmt.Println("My favorite number is", rand.Intn(10))

	// Second Package
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// Only Uppercase is exported (math.pi doesn't work!)
	fmt.Println(math.Pi)

	// First self made, non exported function
	fmt.Println(add(4, 7))

	// Second function, returning two values, Assigning two values from function
	// a, b := equals var a, b =
	// only works inside functions
	a, b := swap("世界", "Hello")
	fmt.Println(a, b)

	// Named returning
	fmt.Println(split(17))
}

//First Function
func add(x, y int) int {
	return x + y
}

//Second function, returning two values
func swap(x, y string) (string, string) {
	return y, x
}

//Named returning
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
