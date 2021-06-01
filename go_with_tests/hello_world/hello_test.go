/*
Writing a test is just like writing a function, with a few rules

It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*/

package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Vedant")
		want := "Hello Vedant !!"
		if got != want {
			t.Errorf("Expected output: %q Recieved output: %q\n", want, got)
		}
	})
	t.Run("saying hello to world if the string is empty", func(t *testing.T) {
		got := hello("")
		want := "Hello World !!"
		if got != want {
			t.Errorf("Expected output: %q Recieved output: %q\n", want, got)
		}
	})
}
