package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	r "reflect"
	"strconv"
	"strings"
)

type DBMaker struct {
	root  string
	types inspect.Types
}

func NewDBMaker(root string, types inspect.Types) *DBMaker {
	return &DBMaker{root: root, types: types}
}

func (ue DBMaker) Compute(data interface{}) (ret ScriptDB, err error) {
	name := r.TypeOf(data).Name()
	if cmdType, ok := ue.types[name]; !ok {
		err = errutil.New("couldnt compute unknown type", name)
	} else {
		db := make(DB)

		key := []string{ue.root}
		impl := strings.Split(*cmdType.Implements, ",")
		db.Add(key, &CommandData{impl[0], name})

		c := CommandVisitor{db, cmdType, key}
		//
		if e := inspect.Inspect(ue.types).Visit(c, data); e != nil {
			err = e
		} else {
			ret = db
		}
	}
	return
}

type Visitor struct {
	db       DB
	baseType *inspect.CommandInfo
	key      []string
}

type CommandVisitor Visitor

type ArrayVisitor struct {
	Visitor
	data *ArrayData
}

func (uc CommandVisitor) NewCommand(p *inspect.ParamInfo, cmdType *inspect.CommandInfo) (inspect.Arguments, error) {
	key := append(uc.key, p.Name)
	uc.db.Add(key, &CommandData{uc.baseType.Name, cmdType.Name})
	return CommandVisitor{uc.db, cmdType, key}, nil
}

func (uc CommandVisitor) NewValue(p *inspect.ParamInfo, v interface{}) (err error) {
	if v != nil {
		uses, _ := p.Usage(false)
		key := append(uc.key, p.Name)
		uc.db.Add(key, &PrimData{uses, v})
	}
	return
}

func (uc CommandVisitor) NewArray(p *inspect.ParamInfo, cmdType *inspect.CommandInfo) (inspect.Elements, error) {
	key := append(uc.key, p.Name)
	return ArrayVisitor{Visitor{uc.db, cmdType, key}, nil}, nil
}

func (ua ArrayVisitor) NewCommand(p *inspect.ParamInfo, cmdType *inspect.CommandInfo) (inspect.Arguments, error) {
	if ua.data == nil {
		ua.data = &ArrayData{cmdType.Name, nil, 1}
		ua.db.Add(ua.key, ua.data)
	}

	kid := strconv.FormatInt(int64(ua.data.Next), 10)
	ua.data.Array = append(ua.data.Array, kid)
	ua.data.Next++

	key := append(ua.key, kid)
	ua.db.Add(key, &CommandData{ua.baseType.Name, cmdType.Name})

	return CommandVisitor{ua.db, cmdType, key}, nil
}
