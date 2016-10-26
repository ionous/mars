package rt

type NumberStream interface {
	HasNext() bool
	GetNext() (Number, error)
}
type TextStream interface {
	HasNext() bool
	GetNext() (Text, error)
}
type ObjectStream interface {
	HasNext() bool
	GetNext() (Object, error)
}
