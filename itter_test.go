package integers

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("ex[ected %q but got %q", expected, repeated)
	}
}
