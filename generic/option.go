package generic

import "errors"

type Option[T any] struct {
	val *T
}

func NewOption[T any]() *Option[T] {
	return &Option[T]{}
}

var ErrOptionIsNone = errors.New("gonads: Option[T] has no value")

func (o *Option[T]) IsSome() bool {
	return o.val != nil
}

func (o *Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o *Option[T]) Take() (T, error) {
	if o.IsNone() {
		var zero T
		return zero, ErrOptionIsNone
	}
	return *o.val, nil
}

func (o *Option[T]) Set(val T) {
	o.val = &val
}

func (o *Option[T]) Clear() {
	o.val = nil
}

func (o *Option[T]) Yank() T {
	if o.IsNone() {
		panic("tip: Yank on None Option")
	}
	return *o.val
}
