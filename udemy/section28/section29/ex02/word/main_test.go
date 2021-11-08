package word

import (
	"ex02/quote"
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count(quote.SunAlso)
	expected := 1349
	if expected != got {
		t.Error("Expected", expected, "Got", got)
	}
}

func ExampleCount() {
	fmt.Println(Count("How are you?"))
	// Output:
	// 3
}
