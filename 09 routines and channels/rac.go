package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(msg)
	}
}

func sum(toSum []int, c chan int) {
	sum := 0
	for _, i := range toSum {
		sum += i
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		time.Sleep(100 * time.Millisecond)
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func main() {
	// Basic routine
	go say("Hello")
	say("世界")

	// Channel Routine
	fmt.Println()
	fmt.Println("## Channel synchronized routine")
	s := []int{7, 2, 8, -9, 4, 0}
	// Channel creation
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// By default, sends and receives block until the other side is ready
	x, y := <-c, <-c //recive from c into two variables
	fmt.Println(x, y, x+y)

	// Range / Close with channels
	fmt.Println()
	fmt.Println("## Channels and close <- fibonacci")
	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("## Select with channels")
	c1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c1)
		}
		quit <- -1 << 12 // Ganz egal was quit zugeordnet wird es geht nur darum ob im quit was reinkommt (siehe switch)
	}()
	fibonacciSelect(c1, quit)

}
