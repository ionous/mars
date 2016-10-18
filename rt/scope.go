package rt

type Value interface{}

type Scope interface {
	FindValue(string) (Value, error)
}

type IndexInfo struct {
	Index           int
	IsFirst, IsLast bool
}
