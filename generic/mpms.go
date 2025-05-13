package generic

type Queue[T any] struct {
	data chan T
}

func NewQueue[T any](size int) Queue[T] {
	return Queue[T]{
		data: make(chan T, size),
	}
}

func (q Queue[T]) Push(v T) {
	q.data <- v
}

func (q Queue[T]) Pop() T {
	return <-q.data
}

func (q Queue[T]) TryPop() (T, bool) {
	select {
	case val := <-q.data:
		return val, true
	default:
		var zero T
		return zero, false
	}
}
