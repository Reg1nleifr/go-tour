package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Object methods via structs
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer methods to change the object itseflf
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Basic Interface
type Abser interface {
	Abs() float64
}

//Advanced Interface Stuff
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	//Graceful nullcheck if the object itself is null / nil
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type Person struct {
	Name string
	Age  int
}

// Stringer Method (ToString())
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// Basic Object Method call
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// Self altering Methods via Pointers
	fmt.Println()
	fmt.Println("## Self altering method call")
	v.Scale(10)
	fmt.Println(v.Abs())

	// Basic interface
	fmt.Println()
	fmt.Println("## Basic interface")
	var a Abser
	a = &v
	fmt.Println(a.Abs())

	// Nil values with interfaces
	fmt.Println()
	fmt.Println("## nil values")
	var in I         //create interface
	var t *T         //create type pointer
	in = t           //set the interface to the pointer
	describe(in)     //describe with null value
	in.M()           //call the method
	in = &T{"hello"} //create value for the interface
	describe(in)
	in.M()

	// Type assertion
	fmt.Println()
	fmt.Println("## Type assertions")
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// f = i.(float64) // panic
	// fmt.Println(f)

	// Type switching
	fmt.Println()
	fmt.Println("## Type switching")
	do(21)
	do("hello")
	do(true)

	// Stringer
	fmt.Println()
	fmt.Println("## Stringer")
	first := Person{"Arthur Dent", 42}
	second := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(first, second)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
