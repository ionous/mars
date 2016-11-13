package rt

type Execute interface {
	Execute(Runtime) error
}
type BoolEval interface {
	GetBool(Runtime) (Bool, error)
}
type NumberEval interface {
	GetNumber(Runtime) (Number, error)
}
type TextEval interface {
	GetText(Runtime) (Text, error)
}
type ObjEval interface {
	GetObject(Runtime) (Object, error)
}
type StateEval interface {
	GetState(Runtime) (State, error)
}
type NumListEval interface {
	GetNumberStream(Runtime) (NumberStream, error)
}
type TextListEval interface {
	GetTextStream(Runtime) (TextStream, error)
}
type ObjListEval interface {
	GetObjStream(Runtime) (ObjectStream, error)
}
