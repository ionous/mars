package rt

type BoolEval interface {
	GetBool(Runtime) (bool, error)
}
type NumEval interface {
	GetNumber(Runtime) (Number, error)
}
type TextEval interface {
	GetText(Runtime) (Text, error)
}
type RefEval interface {
	GetReference(Runtime) (Reference, error)
}
type ListEval interface {
	GetCount() int
}
type NumListEval interface {
	ListEval
	GetNumberIdx(Runtime, int) (Number, error)
}
type TextListEval interface {
	ListEval
	GetTextIdx(Runtime, int) (Text, error)
}
type RefListEval interface {
	ListEval
	GetReferenceIdx(Runtime, int) (Reference, error)
}
