package rt

type Execute interface {
	Execute(Runtime) error
}
type BoolEval interface {
	GetBool(Runtime) (Bool, error)
}
type NumEval interface {
	GetNumber(Runtime) (Number, error)
}
type TextEval interface {
	GetText(Runtime) (Text, error)
}
type ObjEval interface {
	GetObject(Runtime) (Object, error)
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
type ObjListEval interface {
	ListEval
	GetReferenceIdx(Runtime, int) (Reference, error)
}
