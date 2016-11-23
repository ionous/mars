package main

import (
	"encoding/json"
	"fmt"
	"github.com/ionous/mars/script"
	"github.com/ionous/sashimi/util/errutil"
	r "reflect"
	"strings"
)

type TypeParameters struct {
	Name      string  `json:"name"`
	Phrase    *string `json:"phrase,omitempty"`
	Uses      string  `json:"uses"`
	UsesArray *bool   `json:"usesArray,omitempty"`
}

type TypeBlock struct {
	Name       string           `json:"name"`
	Implements *string          `json:"implements,omitempty"`
	Parameters []TypeParameters `json:"params,omitempty"`
}

func AddInterface(f r.StructField) TypeBlock {
	return TypeBlock{Name: f.Name, Implements: NewString("interface")}
}

type Builder struct {
	gen   TypeExists
	types TypeBlocks
	faces Interfaces
}

type TypeExists map[r.Type]bool

func (b *Builder) Add(tb TypeBlock) {
	b.types = append(b.types, tb)
}

func NewString(s string) *string {
	ret := new(string)
	*ret = s
	return ret
}
func NewBool(b bool) *bool {
	ret := new(bool)
	*ret = b
	return ret
}

func (b *Builder) AddParams(s r.Type) (ps []TypeParameters, err error) {
	for i, cnt := 0, s.NumField(); i < cnt; i++ {
		f := s.Field(i)
		tp := TypeParameters{}
		tp.Name = f.Name
		if phrase := f.Tag.Get("mars"); phrase != "" {
			tp.Phrase = NewString(phrase)
		}

		if uses, isArray, e := b.AddParamType(f.Type); e != nil {
			err = errutil.New("error adding", s, "field", f.Name, e)
			break
		} else {
			tp.Uses = uses
			if isArray {
				tp.UsesArray = NewBool(true)
			}
		}
		ps = append(ps, tp)
	}
	return
}

func (b *Builder) AddParamType(s r.Type) (uses string, isArray bool, err error) {
	uses = s.Name()
	switch k := s.Kind(); k {
	case r.String:
		err = b.AddPrim(s)
	case r.Array, r.Slice:
		elem := s.Elem()
		uses, isArray = elem.Name(), true
		if sk := elem.Kind(); sk == r.Struct {
			err = b.AddStruct(elem)
		} else if sk != r.Interface {
			err = errutil.New("array not supported", sk)
		}
	case r.Bool,
		r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64,
		r.Float32, r.Float64:
		{
			err = errutil.New("not supported yet", k)
		}
	case r.Interface:
		if !b.faces.Contains(s) {
			err = errutil.New("has unknown interface", s)
		}
	default:
		err = errutil.New("has unsupported", k)
	}
	return
}
func (b *Builder) AddPrim(s r.Type) (err error) {
	if !b.gen[s] {
		b.gen[s] = true
		name, knd := s.Name(), s.Kind().String()
		var impl *string
		if name != knd {
			impl = NewString(knd)
		}
		tb := TypeBlock{
			Name:       name,
			Implements: impl,
		}
		b.Add(tb)
	}
	return
}

func (b TypeBlocks) Marshall() ([]byte, error) {
	return json.MarshalIndent(b, "", " ")
}

type TypeBlocks []TypeBlock

// meeds a bit of recursion.
func (b *Builder) AddStruct(s r.Type) (err error) {
	if !b.gen[s] {
		if s.Kind() != r.Struct {
			err = errutil.New("not a struct type", s)
		} else {
			b.gen[s] = true
			if face, e := b.faces.FindMatching(s); e != nil {
				err = e
			} else if ps, e := b.AddParams(s); e != nil {
				err = e
			} else {
				tb := TypeBlock{
					Name:       s.Name(),
					Implements: NewString(face),
					Parameters: ps,
				}
				b.Add(tb)
			}
		}
	}
	return
}

type InterfaceRecord struct {
	name string
	face r.Type
}
type Interfaces []InterfaceRecord

func (faces Interfaces) String() string {
	str := make([]string, len(faces))
	for i, s := range faces {
		str[i] = s.name
	}
	return strings.Join(str, ",")
}
func (faces Interfaces) Contains(s r.Type) (okay bool) {
	for _, n := range faces {
		if n.face == s {
			okay = true
			break
		}
	}
	return
}
func (faces Interfaces) FindMatching(s r.Type) (ret string, err error) {
	found := false
	for _, n := range faces {
		u := n.face
		if s.Implements(u) || r.PtrTo(s).Implements(u) {
			if found {
				err = errutil.New("two implementations")
				break
			}
			ret, found = n.name, true
		}
	}
	if !found {
		err = errutil.New("not found", s, faces)
	}
	return
}

func Build() (TypeBlocks, error) {
	var err error
	ref := r.TypeOf((*script.Structures)(nil)).Elem()
	b := Builder{gen: make(TypeExists)}
	for i, fields := 0, ref.NumField(); i < fields; i++ {
		f := ref.Field(i)
		elem := f.Type.Elem()
		if k := elem.Kind(); k == r.Interface {
			tb := AddInterface(f)
			b.Add(tb)
			b.faces = append(b.faces, InterfaceRecord{tb.Name, elem})
			//
		} else if k == r.Struct {
			if e := b.AddStruct(elem); e != nil {
				err = errutil.New("erroring adding structure", i, e)
				break
			}
		} else {
			err = errutil.New("not handled", k)
			break
		}
	}
	return b.types, err
}

func main() {
	if tb, e := Build(); e != nil {
		panic(e)
	} else if m, e := tb.Marshall(); e != nil {
		panic(e)
	} else {
		fmt.Println(string(m))
	}
}

//  "name": "Fragments"
//     "phrase": "[fragment]"
//     "uses": "" <-- blank , should be "Fragment"
//     // "usesArray": true <-=- should be true!
