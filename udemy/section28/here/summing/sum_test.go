package summing

import (
	"fmt"
	"testing"
)

func TestSumming(t *testing.T) {
	got := Summing(1, 2, 3)
	expected := 6
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func ExampleSumming() {
	fmt.Println(Summing(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Output:
	// 55
}

// Summing N times. N is chosed by testing package.
func BenchmarkSumming(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Summing(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}
