package encode

import (
	r "reflect"
)

type DataBlock struct {
	Name string `json:"cmd"`
	Args ArgMap `json:"args,omitempty"`
}

type ArgMap map[string]interface{}
type DataBlocks []*DataBlock

type UniformEncoder struct {
	types TypeMap
}

func NewUniformEncoder(types TypeMap) UniformEncoder {
	return UniformEncoder{types: types}
}

func (ue UniformEncoder) Compute(data interface{}) (ret DataBlock, err error) {
	name := r.TypeOf(data).Name()
	w := NewWalker(ue.types)
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

func (uc UniformCommand) NewCommand(p *ParamInfo, cmdType *CommandType) (ArgWalker, error) {
	dataPtr := &DataBlock{cmdType.Name, make(ArgMap)}
	uc.dataPtr.Args[p.Name] = dataPtr
	return UniformCommand{dataPtr}, nil
}

func (uc UniformCommand) NewValue(p *ParamInfo, v interface{}) (err error) {
	if v != nil {
		uc.dataPtr.Args[p.Name] = v
	}
	return
}

func (uc UniformCommand) NewArray(p *ParamInfo, cmdType *CommandType) (CommandWalker, error) {
	return UniformArray{uc.dataPtr, p.Name}, nil
}

func (ua UniformArray) NewCommand(p *ParamInfo, cmdType *CommandType) (ArgWalker, error) {
	dataPtr := &DataBlock{cmdType.Name, make(ArgMap)}

	array := ua.parent.Args[ua.field]
	if array == nil {
		array = DataBlocks{}
	}
	ua.parent.Args[ua.field] = append(array.(DataBlocks), dataPtr)

	return UniformCommand{dataPtr}, nil
}
