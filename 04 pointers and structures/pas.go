package main

import (
  "fmt"
)

// Basic structure
type Vertex struct {
  X, Y int
}

func main () {
  // Pointers
  i, j := 42, 2701

	p := &i          // point to i
	fmt.Println(*p)  // read i through the pointer
	*p = 27          // set i through the pointer
	fmt.Println(i)   // see the new value of i

	p = &j           // point to j
	*p = *p / 37     // divide j through the pointer
	fmt.Println(j)   // see the new value of j

  // Basic structure
  v := Vertex{5, 6}
  fmt.Println(v)
  v.X = 1
  fmt.Println(v.X)
  // Pointers and structure
  vp := &v
  (*vp).X = 1e9
  fmt.Println(vp.X) //No need to write (*vp) when using structure pointers

  // Allocation of structures
  v1 = Vertex{1, 2}  // has type Vertex
  v2 = Vertex{X: 1}  // Y:0 is implicit
  v3 = Vertex{}      // X:0 and Y:0
  p  = &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, p, v2, v3)
}
