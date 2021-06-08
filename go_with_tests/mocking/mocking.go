package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}
type Spysleeper struct {
	Calls int
}

type DefaultSleeper struct {
	Calls int
}

func (s *Spysleeper) Sleep() {
	s.Calls++
}

func (s *DefaultSleeper) Sleep() {
	s.Calls++
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 3; i >= 1; i-- {
		fmt.Fprint(w, i, "\n")
		sleeper.Sleep()
	}
	fmt.Fprint(w, "Go!")
}

func main() {
	sleeper := DefaultSleeper{
		Calls: 3,
	}
	Countdown(os.Stdout, &sleeper)
}
