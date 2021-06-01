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
	got := hello()
	want := "Hello World !!"

	if got != want {
		// %q is used to print the string with quotes
		t.Errorf("Expected output: %q\nRecieved output: %q\n", want, got)
	}
}
