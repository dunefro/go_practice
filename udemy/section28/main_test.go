package main

import "testing"

func TestMysum(t *testing.T) {
	s := mySum(2, 3)
	expected := 5
	if s != expected {
		t.Error("Expected", expected, "Got", s)
	}
}

func TestTableAll(t *testing.T) {
	type test struct {
		data []int
		sum  int
	}
	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{1, 2, 3, 4, 5, 6, 7}, 28},
		test{[]int{-1, 0, 1}, 0},
		test{[]int{1, 3, 5, 7}, 16},
		test{[]int{-10, -9, -8, 4}, -23},
	}
	for _, v := range tests {
		expected := v.sum
		got := mySum(v.data...)
		if expected != got {
			t.Error("Expected", expected, "got", got)
		}
	}
}
