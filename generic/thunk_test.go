package generic

import "testing"

func TestNewThunkFib(t *testing.T) {
	cache := make([]*Thunk[int], 41)

	var fib func(int) int
	fib = func(i int) int {
		if i == 0 {
			return 0
		}
		if i == 1 || i == 2 {
			return 1
		}
		return cache[i-1].Force() + cache[i-2].Force()
	}

	for i := range cache {
		i := i
		t := NewThunk(func() int { return fib(i) })
		cache[i] = &t
	}

	t.Log(cache[40].Force())
}
