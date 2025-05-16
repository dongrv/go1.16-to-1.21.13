package generic

type Integer interface {
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64 | ~uint | ~int
}

type Iteration[T Integer] interface {
	Iter() T
}

type Iterator[T Integer] struct{ n T }

func (i *Iterator[T]) Iter() T {
	i.n++
	return i.n
}

type Iterator2[T Integer] struct{ doer func() T }

func (i *Iterator2[T]) Iter() T {
	if i.doer != nil {
		return i.doer()
	}
	return 0
}
