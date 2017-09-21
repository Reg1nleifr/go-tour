package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var m2 = make(map[string]int)

func main() {
	// Basic maps
	fmt.Println(m)

	fmt.Println()
	fmt.Println("## Mutating maps")
	// Mutating maps
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])
	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])
	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])
	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println()
	fmt.Println("## functions as variables")
	// Functions can be used as variables and arguments
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
