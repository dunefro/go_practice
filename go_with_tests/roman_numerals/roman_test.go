package main

import "testing"

func TestConvertToRoman(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"Converting 1 to I", 1, "I"},
		{"Converting 1 to II", 2, "II"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("Recieved: %q, Expected: %q", got, test.Want)
			}
		})
	}

}
