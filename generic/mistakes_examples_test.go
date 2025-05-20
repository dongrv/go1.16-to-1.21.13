package generic

import "testing"

func TestNewTypeOK(t *testing.T) {
	a, b := 1, 2
	n := NewTypeOK[*int]{&a, &b}
	t.Logf("%v", n)

	t.Run("Hello", func(t *testing.T) {
		//Hello(1) // Cannot use int as the type INumber3 Type does not implement constraint 'INumber3' because constraint type set is empty
	})
}
