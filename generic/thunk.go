package generic

type Thunk[T any] struct {
	doer func() T
	o    *Option[T]
}

func (t Thunk[T]) Force() T {
	if t.o.IsSome() {
		return t.o.Yank()
	}
	t.o.Set(t.doer())
	return t.o.Yank()
}

func NewThunk[T any](doer func() T) Thunk[T] {
	return Thunk[T]{
		doer: doer,
		o:    NewOption[T](),
	}
}
