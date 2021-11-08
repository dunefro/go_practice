package dog

import "testing"

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
