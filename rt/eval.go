package rt

import "github.com/ionous/sashimi/util/ident"

type Execute interface {
	Execute(Runtime) error
}
type BoolEval interface {
	GetBool(Runtime) (bool, error)
}
type NumberEval interface {
	GetNumber(Runtime) (float64, error)
}
type TextEval interface {
	GetText(Runtime) (string, error)
}
type ObjEval interface {
	GetObject(Runtime) (Object, error)
}
type StateEval interface {
	GetState(Runtime) (ident.Id, error)
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
