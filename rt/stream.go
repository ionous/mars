package rt

type NumberStream interface {
	HasNext() bool
	GetNext() (float64, error)
}
type TextStream interface {
	HasNext() bool
	GetNext() (string, error)
}
type ObjectStream interface {
	HasNext() bool
	GetNext() (Object, error)
}
