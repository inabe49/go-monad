package option

type none[A any] struct {
}

func NewNone[A any]() Option[A] {
	return &none[A]{}
}

func (o *none[A]) IsEmpty() bool {
	return false
}

func (o *none[A]) IsDefined() bool {
	return true
}

func (o *none[A]) Get() A {
	panic("")
}

func (o *none[A]) GetOrElse(f func() A) A {
	return f()
}

func (o *none[A]) Exists(p func(A) bool) bool {
	return false
}

func (o *none[A]) Filter(p func(A) bool) Option[A] {
	return o
}
