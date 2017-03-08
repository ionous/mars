package inspect

import (
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	r "reflect"
	"strconv"
	"strings"
)

type Walker struct {
	types Types
}

func Inspect(types Types) *Walker {
	return &Walker{types}
}

func (w *Walker) Visit(c Arguments, value interface{}) (err error) {
	return w.VisitPath(nil, c, value)
}

func (w *Walker) VisitPath(path Path, c Arguments, value interface{}) (err error) {
	v := r.ValueOf(value)
	name := v.Type().Name()
	if cmdType, ok := w.types[name]; !ok {
		err = errutil.New("type not found", sbuf.Q(name))
	} else {
		err = w.visitArgs(path, c, cmdType, v)
	}
	return
}

func (w *Walker) visitArgs(path Path, c Arguments, cmdType *CommandInfo, cmdData r.Value) (err error) {
	for _, p := range cmdType.Parameters {
		if fieldVal := cmdData.FieldByName(p.Name); !fieldVal.IsValid() {
			err = errutil.New("field not found", sbuf.Q(p.Name))
			break
		} else if fieldVal, e := unpack(fieldVal); e != nil {
			err = e
			break
		} else {
			p := p // pin
			kid := path.ChildPath(p.Name)
			if e := w.visitArg(kid, c, &p, fieldVal); e != nil {
				err = errutil.New("error converting", kid, "because", e)
				break
			}
		}
	}
	return
}

func (w *Walker) visitArray(path Path, c Arguments, baseType *CommandInfo, pv *r.Value) (err error) {
	var cnt int
	if pv != nil {
		cnt = pv.Len()
	}
	if a, e := c.NewArray(path, baseType, cnt); e != nil {
		err = e
	} else if cnt > 0 {
		for i := 0; i < cnt; i++ {
			if elVal, e := unpack(pv.Index(i)); e != nil {
				err = e
				break
			} else {
				kid := path.ChildPath(strconv.Itoa(i))
				if e := w.visitEl(kid, a, baseType, elVal); e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

func (w *Walker) visitEl(kid Path, a Elements, baseType *CommandInfo, pv *r.Value) (err error) {
	if pv == nil {
		_, err = a.NewElement(kid, nil)
	} else {
		if cmdType, e := w.commandType(baseType, *pv); e != nil {
			err = e
		} else if elWalker, e := a.NewElement(kid, cmdType); e != nil {
			err = e
		} else {
			err = w.visitArgs(kid, elWalker, cmdType, *pv)
		}
	}
	return
}

func (w *Walker) visitArg(kid Path, c Arguments, p *ParamInfo, pv *r.Value) (err error) {
	k := pv.Kind()
	u := p.Usage()
	isArray, wantArray := (r.Slice == k), u.IsArray()

	if isArray != wantArray {
		err = errutil.New("array mismatch")
	} else {
		// commands start uppercase, primitives lowercase.
		if uses := u.Uses(); u.IsCommand() {
			if baseType, ok := w.types[uses]; !ok {
				err = errutil.New("type not found", uses)
			} else {
				if !isArray {
					err = w.visitCmd(kid, c, baseType, pv)
				} else {
					err = w.visitArray(kid, c, baseType, pv)
				}
			}
		} else {
			if !isArray {
				if prim, e := convertPrim(uses, *pv); e != nil {
					err = e
				} else {
					err = c.NewValue(kid, prim)
				}
			} else {
				if prim, e := convertArray(uses, *pv); e != nil {
					err = e
				} else {
					err = c.NewValue(kid, prim)
				}
			}
		}
	}
	return
}

func (w *Walker) visitCmd(kid Path, c Arguments, baseType *CommandInfo, pv *r.Value) (err error) {
	if pv == nil {
		_, err = c.NewCommand(kid, baseType, nil)
	} else {
		if cmdType, e := w.commandType(baseType, *pv); e != nil {
			err = e
		} else if elWalker, e := c.NewCommand(kid, baseType, cmdType); e != nil {
			err = e
		} else {
			err = w.visitArgs(kid, elWalker, cmdType, *pv)
		}
	}
	return
}

// can return nil
func (w *Walker) commandType(baseType *CommandInfo, v r.Value) (ret *CommandInfo, err error) {
	name := v.Type().Name()
	if cmdType, ok := w.types[name]; !ok {
		err = errutil.New("type not found", name)
	} else if !implements(baseType, cmdType) {
		err = errutil.New("expected implementor of", baseType.Name, "got", cmdType.Name, *cmdType.Implements)
	} else {
		ret = cmdType
	}
	return
}

func unpack(v r.Value) (ret *r.Value, err error) {
	switch k := v.Kind(); k {
	default:
		ret = &v
	case r.Ptr, r.Interface:
		if !v.IsNil() {
			ret, err = unpack(v.Elem())
		}
		// default:
		// 	err = errutil.New("arg not supported", sbuf.Q(k))
	}
	return
}

// does cmd implement base
func implements(base, cmd *CommandInfo) (okay bool) {
	if cmd.Implements != nil {
		for _, k := range strings.Split(*cmd.Implements, ",") {
			if k == base.Name {
				okay = true
				break
			}
		}
	}
	return
}

// check the passed value is of the expected primitive type
func convertPrim(uses string, val r.Value) (ret interface{}, err error) {
	prim := map[r.Kind]string{
		r.String:  "string",
		r.Bool:    "bool",
		r.Float64: "float64",
	}
	kind := val.Kind()
	if t, ok := prim[kind]; (!ok && uses != "blob") || (ok && t != uses) {
		err = errutil.New("primitive type mismatch", uses, kind.String())
	} else {
		ret = val.Interface()
	}
	return
}

// change to an object containing a slice of raw interface values
func convertArray(uses string, v r.Value) (ret interface{}, err error) {
	if cnt := v.Len(); cnt > 0 {
		var ar []interface{}
		for i := 0; i < cnt; i++ {
			if prim, e := convertPrim(uses, v.Index(i)); e != nil {
				err = errutil.New("error converting array of", uses, "at", i, "because", e)
				break
			} else {
				ar = append(ar, prim)
			}
		}
		if err == nil {
			ret = ar
		}
	}
	return
}
