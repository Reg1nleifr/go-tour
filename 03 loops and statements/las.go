package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	//for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// "while"
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// ifs and else
	fmt.Println(sqrt(2), sqrt(-4))

	//if/else with statment
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//switch case: breaks automatically
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
		fallthrough
	default:
		fmt.Println("Too far away.")
	}

	//defer --> Executes Statement at return of function
	defer fmt.Printf("世界")
	fmt.Printf("Hello, ")

	//defer can be stacked --> Last in first out
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d.. ", i)
	}

}

// ifs and else
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if/else with statment
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
