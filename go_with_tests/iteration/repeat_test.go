package iteration

import "testing"

func TestRepeat(t *testing.T) {

	checkgotandwant := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Expected Output: %q Received Output %q", got, want)
		}
	}
	t.Run("Testing without 0 no of iterations", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaaaa"
		checkgotandwant(t, got, want)
	})
	t.Run("Testing with fixed no of iterations", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"
		checkgotandwant(t, got, want)
	})
}

// Benchmarking test
// here we use testing.B
// it will run for b.N times and call the function Repeat
// To run this test we execute the command "go test -bench=. "
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
