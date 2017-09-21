package main

import (
	"fmt"
)

// Splice #2
var names = [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}

type Vertex struct {
	X, Y float64
}

func main() {
	// Basic Array
	var a [2]string
	a[1] = "World"
	a[0] = "世界" //Shíjiè
	fmt.Println(a[1], a[0])
	fmt.Println(a)

	// Splice
	primes := [6]int{2, 3, 5, 7, 11, 13}
	x := primes[0:2]
	fmt.Println(x)

	// Splice #2 is like a reference
	fmt.Println(primes)
	// Create a second splice
	y := primes[1:3]
	fmt.Println(x, y)
	// Change a value contained in both splices
	y[0] = 666
	fmt.Println(x, y)
	fmt.Println(primes)

	// Array literals
	s := []struct {
		i int
		b bool
	}{
		{2, true}, {3, false},
		{5, true}, {7, true},
		{11, false}, {13, true},
	}
	fmt.Println(s)

	// Slice defaults
	// these slice expressions are equivalent:
	sd := primes[0:6]
	sd = primes[:6]
	sd = primes[0:]
	sd = primes[:]
	fmt.Println(sd)

	fmt.Println()
	fmt.Println("## Slicing around")
	// Length and capacity
	ints := []int{2, 3, 5, 7, 11, 13}
	printSlice(ints)
	// Slice the slice to give it zero length.
	ints = ints[:0]
	printSlice(ints)
	// Extend its length.
	ints = ints[:4]
	printSlice(ints)
	// Drop its first two values.
	ints = ints[2:]
	printSlice(ints)
	// Extend its length out of range.
	// ints = ints[:8]
	printSlice(ints)

	fmt.Println()
	fmt.Println("## Dynamic Array")
	// Dynamic Arrays (only works when decreasing the front -> decreases cap)
	b := make([]int, 5)
	printSlice(b)
	b = b[2:]
	printSlice(b)

	fmt.Println()
	fmt.Println("## append elements")
	// Append elements -> Seems to work exponentionally on the size before
	// append works on nil slices.
	b = append(b, 0)
	printSlice(b)
	// We can add more than one element at a time.
	b = append(b, 2, 3, 4)
	printSlice(b)

	// for each
	fmt.Println()
	fmt.Println("## for range")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	fmt.Println()
	fmt.Println("## for range escaping index or value")
	// You can skip the index or value by assigning to _.
	// If you only want the index, drop the ", value" entirely.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
