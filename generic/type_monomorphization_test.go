package generic

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, Search[int]([]int{1, 2, 3}, 2), true)
	assert.Equal(t, Search[float32]([]float32{1.1, 2.2, 3.3}, 1), false)
	assert.Equal(t, Search[string]([]string{"1", "2", "3"}, "2"), true)
}
