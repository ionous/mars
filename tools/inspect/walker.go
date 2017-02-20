package inspect

import (
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	r "reflect"
	"strings"
)

type Walker struct {
	types Types
}

func Inspect(types Types) *Walker {
	return &Walker{types}
}

func (w *Walker) Visit(c Arguments, value interface{}) (err error) {
	v := r.ValueOf(value)
	name := v.Type().Name()
	if cmdType, ok := w.types[name]; !ok {
		err = errutil.New("type not found", sbuf.Q(name))
	} else {
		err = w.visitArgs(c, cmdType, v)
	}
	return
}

func (w *Walker) visitArgs(c Arguments, cmdType *CommandInfo, cmdData r.Value) (err error) {
	for _, p := range cmdType.Parameters {
		if fieldVal := cmdData.FieldByName(p.Name); !fieldVal.IsValid() {
			err = errutil.New("field not found", sbuf.Q(p.Name))
			break
		} else if e := w.visitArg(c, &p, fieldVal); e != nil {
			err = errutil.New("error converting", cmdType.Name+"."+p.Name, "because", e)
			break
		}
	}
	return
}

func (w *Walker) visitArray(c Arguments, p *ParamInfo, baseType *CommandInfo, v r.Value) (err error) {
	if a, e := c.NewArray(p, baseType); e != nil {
		err = e
	} else {
		for i := 0; i < v.Len(); i++ {
			elVal := v.Index(i)
			if e := w.visitCmd(a, p, baseType, elVal); e != nil {
				err = e
				break
			}
		}
	}
	return
}

func (w *Walker) visitCmd(cw Elements, p *ParamInfo, baseType *CommandInfo, v r.Value) (err error) {
	switch k := v.Kind(); k {
	case r.Struct:
		if cmdType, e := w.commandType(v); e != nil {
			err = e
		} else if !implements(baseType, cmdType) {
			err = errutil.New("expected implementor of", baseType.Name, "got", cmdType.Name, *cmdType.Implements)
		} else {
			if elWalker, e := cw.NewCommand(p, cmdType); e != nil {
				err = e
			} else {
				err = w.visitArgs(elWalker, cmdType, v)
			}
		}
	case r.Ptr, r.Interface:
		if !v.IsNil() {
			err = w.visitCmd(cw, p, baseType, v.Elem())
		}
	default:
		err = errutil.New("arg not supported", sbuf.Q(k))
	}
	return
}

func (w *Walker) visitArg(c Arguments, p *ParamInfo, v r.Value) (err error) {
	k := v.Kind()
	uses, style := p.Usage(true)
	isArray, wantArray := (r.Slice == k), style["array"] == "true"

	if isArray != wantArray {
		err = errutil.New("array mismatch")
	} else {
		// commands are uppercase, primitives lowercase.
		if strings.ToUpper(uses[:1]) == uses[:1] {
			if cmdType, ok := w.types[uses]; !ok {
				err = errutil.New("type not found", cmdType.Name)
			} else {
				if !isArray {
					err = w.visitCmd(c, p, cmdType, v)
				} else {
					err = w.visitArray(c, p, cmdType, v)
				}
			}
		} else {
			if !isArray {
				if prim, e := convertPrim(uses, v); e != nil {
					err = e
				} else {
					err = c.NewValue(p, prim)
				}
			} else {
				if prim, e := convertArray(uses, v); e != nil {
					err = e
				} else if len(prim) > 0 {
					err = c.NewValue(p, prim)
				}
			}
		}
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

func (w *Walker) commandType(v r.Value) (ret *CommandInfo, err error) {
	t := v.Type()
	name := t.Name()
	if cmd, ok := w.types[name]; !ok {
		err = errutil.New("type not found", name)
	} else {
		ret = cmd
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

// change to a slice of raw interface values
func convertArray(uses string, v r.Value) (ret []interface{}, err error) {
	for i := 0; i < v.Len(); i++ {
		if prim, e := convertPrim(uses, v.Index(i)); e != nil {
			err = errutil.New("error converting array of", uses, "at", i, "because", e)
			break
		} else {
			ret = append(ret, prim)
		}
	}
	return
}
