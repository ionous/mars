package blocks

type CommandData struct {
	Type string
	Cmd  string
}

type PrimData struct {
	Type  string
	Value interface{}
}
type ArrayData struct {
	Type  string
	Array []string
	Next  int
}
