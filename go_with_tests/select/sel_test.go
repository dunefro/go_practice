package sel

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)
	if want != got {
		t.Errorf("We expected %q to be fast but %q was faster", want, got)
	}
}
