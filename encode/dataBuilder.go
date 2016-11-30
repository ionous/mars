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

func ComputeArg(v r.Value) (ret interface{}, err error) {
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
				if l, e := ComputeArg(v.Index(i)); e != nil {
					err = errutil.New("error converting", k, "at", i, "because", e)
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
			if r, e := ComputeArg(v.Elem()); e != nil {
				err = errutil.New("error converting", k, "because", e)
			} else {
				ret = r
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
		ret, err = ComputeCmd(r.ValueOf(src))
	}
	return
}

func ComputeCmd(src r.Value) (ret DataBlock, err error) {
	srcType, args := src.Type(), make(ArgMap)
	if srcType.Kind() != r.Struct {
		err = errutil.New("error", srcType, "is not a struct")
	} else {
		ret.Name = srcType.Name()
		ret.Args = args
		for i := 0; i < srcType.NumField(); i++ {
			f, v := srcType.Field(i), src.Field(i)
			if arg, e := ComputeArg(v); e != nil {
				err = errutil.New("error converting", srcType, "at", f.Name, "because", e)
			} else if arg != nil {
				args[f.Name] = arg
			}
		}
	}
	return
}
