package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

// Decleration of an error type
type ErrNegativeSqrt float64

// error function
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt with error function
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(-2)
	}
	return math.Sqrt(x), nil
}

func main() {
	// Error presentation
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	// Reader
	fmt.Println()
	fmt.Println("## Reader sample")
	stringReader := strings.NewReader("Hello, Reader sample!")

	bytesRead := make([]byte, 8)
	for {
		numberOfBytesRead, err := stringReader.Read(bytesRead)
		fmt.Printf("numberOfBytesRead = %v, err = %v, bytesRead = %v\n", numberOfBytesRead, err, bytesRead)
		fmt.Printf("bytesRead[:numberOfBytesRead] = %q\n", bytesRead[:numberOfBytesRead])
		if err == io.EOF {
			break
		}
	}

}
