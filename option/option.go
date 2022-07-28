package option

type Option[A any] interface {
	IsDefined() bool
	IsEmpty() bool
	Get() A
	GetOrElse(f func() A) A

	Exists(p func(A) bool) bool
	Filter(p func(A) bool) Option[A]
}
