package generic

import "testing"

func TestNewTypeOK(t *testing.T) {
	a, b := 1, 2
	n := NewTypeOK[*int]{&a, &b}
	t.Logf("%v", n)
}
