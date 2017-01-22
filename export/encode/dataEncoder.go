package encode

import (
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	r "reflect"
)

type ArgMap map[string]interface{}

type DataBlock struct {
	Name string `json:"cmd"`
	Args ArgMap `json:"args,omitempty"`
}

type Encoder interface {
	Encode() (DataBlock, error)
}

func (args ArgMap) ComputeArgs(at r.Type, av r.Value) (err error) {
	for i := 0; i < at.NumField(); i++ {
		ft, fv := at.Field(i), av.Field(i)
		// only parse public fields
		if ft.PkgPath == "" {
			// keep adding to this structure as if all of the fields were directly embedded.
			if ft.Anonymous {
				if e := args.ComputeArgs(ft.Type, fv); e != nil {
					err = e
					break
				}
			} else if arg, e := args.ComputeArg(fv); e != nil {
				err = errutil.New("error converting", fv, "at", ft.Name, "because", e)
				break
			} else if arg != nil {
				args[ft.Name] = arg
			}
		}
	}
	return
}

func (args ArgMap) ComputeArg(v r.Value) (ret interface{}, err error) {
	switch k := v.Kind(); k {
	case r.String:
		if s := v.String(); s != "" {
			ret = s
		}
	case r.Bool:
		ret = v.Bool()
	case r.Float64:
		ret = v.Float()

	case r.Slice:
		if cnt := v.Len(); cnt > 0 {
			var out []interface{}
			for i := 0; i < cnt; i++ {
				if l, e := args.ComputeArg(v.Index(i)); e != nil {
					err = errutil.New("error converting", k, "at", i, "because", e)
					break
				} else {
					out = append(out, l)
				}
			}
			ret = out
		}

	case r.Struct:
		if r, e := ComputeCmd(v); e != nil {
			err = errutil.New("error converting", k, "because", e)
		} else {
			ret = r
		}

	case r.Ptr, r.Interface:
		if !v.IsNil() {
			if m, ok := v.Interface().(Encoder); ok {
				if data, e := m.Encode(); e != nil {
					err = e
				} else {
					ret = data
				}
			} else {
				if r, e := args.ComputeArg(v.Elem()); e != nil {
					err = errutil.New("error converting", k, "because", e)
				} else {
					ret = r
				}
			}
		}
		// interesingly: interface storage stores an interface object
		// which contains the underlying object.
	default:
		err = errutil.New("arg not supported", sbuf.Q(k))
	}
	return
}

func Compute(src interface{}) (ret DataBlock, err error) {
	if src != nil {
		v := r.ValueOf(src)
		if v.Kind() == r.Ptr {
			v = v.Elem()
		}

		if r, e := ComputeCmd(v); e != nil {
			err = errutil.New("compute", e)
		} else {
			ret = r
		}
	}
	return
}

func ComputeCmd(src r.Value) (ret DataBlock, err error) {
	args := make(ArgMap)
	srcType := src.Type()
	if srcType.Kind() != r.Struct {
		err = errutil.New(srcType, "is not a struct")
	} else if e := args.ComputeArgs(srcType, src); e != nil {
		err = e
	} else {
		ret.Name = srcType.Name()
		ret.Args = args
	}
	return
}
