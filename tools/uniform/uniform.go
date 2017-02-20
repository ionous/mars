package uniform

import (
	"github.com/ionous/mars/tools/inspect"
	r "reflect"
)

type DataBlock struct {
	Name string `json:"cmd"`
	Args ArgMap `json:"args,omitempty"`
}

type ArgMap map[string]interface{}
type DataBlocks []*DataBlock

type UniformEncoder struct {
	types inspect.Type
}

func NewUniformEncoder(types inspect.Type) UniformEncoder {
	return UniformEncoder{types: types}
}

func (ue UniformEncoder) Compute(data interface{}) (ret DataBlock, err error) {
	name := r.TypeOf(data).Name()
	w := inspect.Inspect(ue.types)
	dataPtr := &DataBlock{name, make(ArgMap)}
	c := UniformCommand{dataPtr}
	if e := w.Visit(c, data); e != nil {
		err = e
	} else {
		ret = *dataPtr
	}
	return
}

type UniformArray struct {
	parent *DataBlock
	field  string
}

type UniformCommand struct {
	dataPtr *DataBlock
}

func (uc UniformCommand) NewCommand(p *inspect.ParamInfo, cmdType *inspect.CommandInfo) (inspect.Arguments, error) {
	dataPtr := &DataBlock{cmdType.Name, make(ArgMap)}
	uc.dataPtr.Args[p.Name] = dataPtr
	return UniformCommand{dataPtr}, nil
}

func (uc UniformCommand) NewValue(p *inspect.ParamInfo, v interface{}) (err error) {
	if v != nil {
		uc.dataPtr.Args[p.Name] = v
	}
	return
}

func (uc UniformCommand) NewArray(p *inspect.ParamInfo, cmdType *inspect.CommandInfo) (inspect.Elements, error) {
	return UniformArray{uc.dataPtr, p.Name}, nil
}

func (ua UniformArray) NewCommand(p *inspect.ParamInfo, cmdType *inspect.CommandInfo) (inspect.Arguments, error) {
	dataPtr := &DataBlock{cmdType.Name, make(ArgMap)}

	array := ua.parent.Args[ua.field]
	if array == nil {
		array = DataBlocks{}
	}
	ua.parent.Args[ua.field] = append(array.(DataBlocks), dataPtr)

	return UniformCommand{dataPtr}, nil
}
