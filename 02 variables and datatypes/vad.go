package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

//Initialization
var c, python, java bool

//Initialization #2
var x, y int = 1, 2

// Data types, bit shifting, complex numbers
var (
  ToBe    bool        = false
  MaxInt  uint64      = 1<<64 - 1
  z       complex128  = cmplx.Sqrt(-5 + 12i)
)

// Default / Zero values
var (
  j int
  f float64
  b bool
  s string
)

func main() {

  // Initialization
  var i int
  fmt.Println(i, c, python, java)

  // Initialization #2
  var csharp, golang, assembler = false, true, "no!"
  fmt.Println(x, y, csharp, golang, assembler)

  // Data types, bit shifting, complex numbers
  fmt.Printf("Type: %T, value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T, value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T, value: %v\n", z, z)

  // Default / Zero values
  fmt.Printf("%v %v %v %q\n", j, f, b, s)

  // Type conversion
  var k, l int = 3, 4
	var f float64 = math.Sqrt(float64(k*k + l*l))
	var c uint = uint(f)
	fmt.Println(k, l, c)

  // Numeric conversion
  var integer  int = 42
  float := float64(i)
  unint := uint(f)
  fmt.Printf("Type: %T, value: %v\n", integer, integer)
  fmt.Printf("Type: %T, value: %v\n", float, float)
  fmt.Printf("Type: %T, value: %v\n", unint, unint)

  //Constants
  const World = "世界"
  fmt.Println("Hello", World)
}
