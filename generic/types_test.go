package generic

import "testing"

func TestTypes(t *testing.T) {
	t.Run("print-generic-type", func(t *testing.T) {
		t.Logf("generic struct: %v\n", s)
		t.Logf("generic group struct: %v\n", gs)
		t.Logf("generic map: %v\n", m)

		c <- 1
		t.Logf("generic chan: %v\n", <-c)

		(Example[int]{}).Print(1)
	})

	t.Run("basic-interface", func(t *testing.T) { LoopBasicInterface() })

	t.Run("hi", func(t *testing.T) {
		person := Person{ID: 1, Name: "Alice"}
		Hi(person)

		american := American{ID: 1, Name: "Alice"}
		Hi(american)

	})
}
