package generic

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestIterator_Iter(t *testing.T) {
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
}
