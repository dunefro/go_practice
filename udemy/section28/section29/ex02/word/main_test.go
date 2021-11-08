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
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
