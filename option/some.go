package option

type some[A any] struct {
	value A
}

func NewSome[A any](value A) Option[A] {
	return &some[A]{
		value: value,
	}
}

func (o *some[A]) IsEmpty() bool {
	return false
}

func (o *some[A]) IsDefined() bool {
	return true
}

func (o *some[A]) Get() A {
	return o.value
}

func (o *some[A]) GetOrElse(f func() A) A {
	return o.value
}

func (o *some[A]) Exists(p func(A) bool) bool {
	return p(o.value)
}

func (o *some[A]) Filter(p func(A) bool) Option[A] {
	if p(o.value) {
		return o
	}

	panic("")
}
