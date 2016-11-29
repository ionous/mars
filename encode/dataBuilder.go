package encode

import (
	"github.com/ionous/sashimi/util/errutil"
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
		ret = v.String()

	case r.Slice:
		var out []interface{}
		for i, cnt := 0, v.Len(); i < cnt; i++ {
			if l, e := ComputeArg(v.Index(i)); e != nil {
				err = e
			} else {
				out = append(out, l)
			}
		}
		ret = out

	case r.Bool,
		r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64,
		r.Float32, r.Float64:
		{
			err = errutil.New("not supported yet", k)
			// panic(err)
		}
	case r.Struct:
		ret, err = ComputeCmd(v)

	case r.Ptr, r.Interface:
		ret, err = ComputeArg(v.Elem())
		// interesingly: interface storage stores an interface object
		// which contains the underlying object.
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
				err = e
			} else {
				args[f.Name] = arg
			}
		}
	}
	return
}
