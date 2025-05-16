package generic

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestIterator_Iter(t *testing.T) {
	// iterator type
	t.Run("int8", func(t *testing.T) {
		iter := Iterator[int8]{}
		for i := 0; i < 99; i++ {
			iter.Iter()
		}
		assert.IsEqual(iter.Iter(), 100)
	})
	t.Run("int32", func(t *testing.T) {
		iter := Iterator[int32]{}
		for i := 0; i < 999; i++ {
			iter.Iter()
		}
		assert.IsEqual(iter.Iter(), 1000)
	})

	// iterator anonymous functions
	t.Run("int8-func", func(t *testing.T) {
		var i int
		iter := Iterator2[int]{doer: func() int {
			i++
			return i
		}}
		for i := 0; i < 9; i++ {
			iter.Iter()
		}
		v := iter.Iter()
		assert.Equal(t, v, 10)
	})
}
