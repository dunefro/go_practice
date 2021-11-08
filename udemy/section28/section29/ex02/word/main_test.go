package word

import (
	"ex02/quote"
	"fmt"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count(quote.SunAlso)
	expected := 1349
	if expected != got {
		t.Error("Expected", expected, "Got", got)
	}
}
func TestUseCount(t *testing.T) {
	s := "How are you you you how How"
	gotm := UseCount(s)
	exm := map[string]int{
		"How": 2,
		"are": 1,
		"you": 3,
		"how": 1,
	}
	if !reflect.DeepEqual(gotm, exm) {
		t.Error("Expected map", exm, "got map", gotm)
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
