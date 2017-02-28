package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	r "reflect"
	"strconv"
	"strings"
)

type DBMaker struct {
	root  inspect.Path
	types inspect.Types
}

func NewDBMaker(root string, types inspect.Types) *DBMaker {
	path := inspect.NewPath(root)
	return &DBMaker{root: path, types: types}
}

func (ue DBMaker) Compute(data interface{}) (ret ScriptDB, err error) {
	name := r.TypeOf(data).Name()
	if cmdType, ok := ue.types[name]; !ok {
		err = errutil.New("couldnt compute unknown type", name)
	} else {
		db := make(DB)

		impl := strings.Split(*cmdType.Implements, ",")
		db.Add(ue.root, &CommandData{impl[0], name})

		c := &CommandVisitor{db, cmdType}
		if e := inspect.Inspect(ue.types).VisitPath(ue.root, c, data); e != nil {
			err = e
		} else {
			ret = db
		}
	}
	return
}

type Visitor struct {
	db            DB
	containerType *inspect.CommandInfo
}

type CommandVisitor Visitor

type ArrayVisitor struct {
	Visitor
	data *ArrayData
}

func (uc *CommandVisitor) NewCommand(path inspect.Path, baseType, cmdType *inspect.CommandInfo) (inspect.Arguments, error) {
	if cmdType != nil {
		uc.db.Add(path, &CommandData{baseType.Name, cmdType.Name})
	}
	return &CommandVisitor{uc.db, cmdType}, nil
}

func (uc *CommandVisitor) NewValue(path inspect.Path, v interface{}) (err error) {
	if p, ok := uc.containerType.FindParam(path.Last()); !ok {
		err = errutil.New("couldnt find parameter for", path)
	} else if v != nil {
		uses, _ := p.Usage(false)
		uc.db.Add(path, &PrimData{uses, v})
	}
	return
}

func (uc *CommandVisitor) NewArray(path inspect.Path, baseType *inspect.CommandInfo, _ int) (inspect.Elements, error) {
	return &ArrayVisitor{Visitor{uc.db, baseType}, nil}, nil
}

func (ua *ArrayVisitor) NewElement(path inspect.Path, cmdType *inspect.CommandInfo) (inspect.Arguments, error) {
	if ua.data == nil {
		ua.data = &ArrayData{ua.containerType.Name, nil, 1}
		ua.db.Add(path, ua.data)
	}

	kid := path.Last()
	if i, e := strconv.Atoi(kid); e == nil {
		if ua.data.Next <= i {
			ua.data.Next = i + 1
		}
	}
	ua.data.Array = append(ua.data.Array, kid)
	ua.db.Add(path, &CommandData{ua.containerType.Name, cmdType.Name})

	return &CommandVisitor{ua.db, cmdType}, nil
}
