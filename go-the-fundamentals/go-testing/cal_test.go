package cal

import "testing"

func TestSumOfFirst(t *testing.T) {
	given := 3
	want := 6

	get := SumOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but got %d", given, want, get)
	}
}
