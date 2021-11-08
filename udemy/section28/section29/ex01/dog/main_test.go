package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	dogAge := 5
	got := Years(dogAge)
	expected := 35
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func TestYearsTwo(t *testing.T) {
	dogAge := 5
	got := YearsTwo(dogAge)
	expected := 35

	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}
