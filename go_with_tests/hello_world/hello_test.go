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

	// function to remove the duplicate codes
	checkgotandwant := func(t testing.TB, got string, want string) {
		// Helper() function tells the compiler that if there is an error in the code and then its not coming from this function
		// It helps developer understand that this function is not causing the actual problem
		// this function is just a helper function
		// If any error comes up and we use Helper() then the error will be shown coming from the place where checkgotandwant() function is called
		// If Helper() function is not used then it will show the error coming from the function below
		t.Helper()
		if got != want {
			t.Errorf("Expected output: %q Recieved output: %q\n", want, got)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Vedant", "")
		want := "Hello Vedant !!"
		checkgotandwant(t, got, want)
	})
	t.Run("saying hello to world if the string is empty", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World !!"
		checkgotandwant(t, got, want)
	})
	// language for greeting in hindi
	t.Run("In Hindi", func(t *testing.T) {
		got := Hello("Vedant", "hindi")
		want := "Namaste Vedant !!"
		checkgotandwant(t, got, want)
	})

	// language for greeting in spanish
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Vedant", "spanish")
		want := "Hola Vedant !!"
		checkgotandwant(t, got, want)
	})
}
