package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("Expected: '%d' Received: '%d' ", got, want)
	}
}

// Below is an example function that is executed
// It is very important to add the comments "Ouput: 6" line in the end which will decide if this function will be executed or not
// To run this execute go test -v
func ExampleAdder() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}
