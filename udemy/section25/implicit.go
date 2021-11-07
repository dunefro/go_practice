package main

import (
	"errors"
	"fmt"
	"math"
)

type MyType struct {
	val int
	err error
}

// Implementing the error interface
func (m MyType) Error() string {
	return fmt.Sprintf("Operation failed with the error %q. Value %d should be greater then 0", m.err.Error(), m.val)
}

func main() {
	// Creating the error
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func sqrt(n int) (int, error) {
	if n < 0 {
		// creating an error
		newerr := errors.New("This is the error field that will be part of struct")
		// return MyType as it implements the error inteface
		return 0, MyType{
			val: n,
			err: newerr,
		}
	}
	return int(math.Sqrt(float64(n))), nil
}
