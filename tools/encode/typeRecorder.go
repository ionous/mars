package encode

import (
	"github.com/ionous/mars"
	"github.com/ionous/sashimi/util/errutil"
	"net/url"
	r "reflect"
	"strings"
)

type TypeRecorder struct {
	allFaces Interfaces
	gen      TypeExists
}

func (b *TypeRecorder) addTypes(p *mars.Package) (ret []TypeBlock, err error) {
	if b.gen == nil {
		b.gen = make(TypeExists)
	}

	if faceTypes, e := b.addInterfaces(p, p.Interfaces); e != nil {
		err = e
	} else if cmdTypes, e := b.addCommands(p, p.Commands); e != nil {
		err = e
	} else {
		ret = append(faceTypes, cmdTypes...)
	}
	return
}

type TypeParameters struct {
	Name   string  `json:"name"`
	Phrase *string `json:"phrase,omitempty"`
	Uses   string  `json:"uses"`
}

type TypeBlock struct {
	Name       string           `json:"name"`
	Implements *string          `json:"implements,omitempty"`
	Parameters []TypeParameters `json:"params,omitempty"`
}

type TypeExists map[r.Type]bool

func newString(s string) (ret *string) {
	if s != "" {
		ret = new(string)
		*ret = s
	}
	return ret
}
func newBool(b bool) *bool {
	ret := new(bool)
	*ret = b
	return ret
}

type Parameters struct {
	ps []TypeParameters
}

func (b *TypeRecorder) addParams(p *mars.Package, s r.Type, ps *Parameters) (err error) {
	if s.Kind() != r.Struct {
		err = errutil.New("couldn't add params of", s)
	} else {
		for i, cnt := 0, s.NumField(); i < cnt; i++ {
			f := s.Field(i)
			tp := TypeParameters{}
			tp.Name = f.Name
			var primType string
			if tags := f.Tag.Get("mars"); tags != "" {
				phraseType := strings.Split(tags, ";")
				tp.Phrase = newString(phraseType[0])
				if cnt := len(phraseType); cnt == 2 {
					primType = phraseType[1]
				}
			}
			//
			if f.Anonymous {
				if e := b.addParams(p, f.Type, ps); e != nil {
					err = e
					break
				}
			} else {
				kinds := make(url.Values)
				if uses, e := b.addParam(p, f.Type, kinds); e != nil {
					err = errutil.New("couldn't add field", f.Name, e)
					break
				} else {
					if primType != "" {
						kinds.Set("type", primType)
					}
					if len(kinds) != 0 {
						uses = uses + "?" + kinds.Encode()

					}
					tp.Uses = uses
					ps.ps = append(ps.ps, tp)
				}
			}
		}
	}
	return
}

func (b *TypeRecorder) addParam(p *mars.Package, s r.Type, kinds url.Values) (uses string, err error) {
	switch n, k := s.Name(), s.Kind(); k {
	case r.String, r.Bool, r.Float64:
		uses = k.String()
		if n != uses {
			kinds.Add("type", n)
		}

	case r.Array, r.Slice:
		uses, err = b.addParam(p, s.Elem(), kinds)
		kinds.Add("array", "true")

	case r.Interface:
		// FIX: for now.
		if n == "Generic" {
			uses = "ObjEval"
		} else if b.allFaces.Contains(s) {
			uses = n
		} else {
			err = errutil.New("has unknown interface", n)
		}
	default:
		err = errutil.New("has unsupported", k)
	}
	return
}

// meeds a bit of recursion.
func (b *TypeRecorder) addStruct(p *mars.Package, s r.Type) (ret *TypeBlock, err error) {
	if !b.gen[s] {
		if s.Kind() != r.Struct {
			err = errutil.New("not a struct type", s)
		} else {
			b.gen[s] = true
			if face, e := b.allFaces.FindMatching(s); e != nil {
				err = e
			} else {
				ps := Parameters{}
				if e := b.addParams(p, s, &ps); e != nil {
					err = e
				} else {
					ret = &TypeBlock{
						Name:       s.Name(),
						Implements: newString(face),
						Parameters: ps.ps,
					}
				}
			}
		}
	}
	return
}

func (b *TypeRecorder) addInterface(p *mars.Package, t r.Type) (ret *TypeBlock, err error) {
	if !b.gen[t] {
		b.gen[t] = true
		name := t.Name()
		b.allFaces = append(b.allFaces, InterfaceRecord{name, t})
		ret = &TypeBlock{
			Name:       name,
			Implements: newString("interface"),
		}
	}
	return
}

func (b *TypeRecorder) addCommands(p *mars.Package, cmds interface{}) (ret []TypeBlock, err error) {
	if cmds != nil {
		ref := r.TypeOf(cmds).Elem()
		for i, fields := 0, ref.NumField(); i < fields; i++ {
			f := ref.Field(i)
			elem := f.Type.Elem()
			if newType, e := b.addStruct(p, elem); e != nil {
				err = errutil.New("error adding command", f.Name, e)
				break
			} else if newType != nil {
				ret = append(ret, *newType)
			}
		}
	}
	return
}

func (b *TypeRecorder) addInterfaces(p *mars.Package, allFaces interface{}) (ret []TypeBlock, err error) {
	if allFaces != nil {
		ref := r.TypeOf(allFaces).Elem()
		for i, fields := 0, ref.NumField(); i < fields; i++ {
			f := ref.Field(i)
			if newType, e := b.addInterface(p, f.Type); e != nil {
				err = errutil.New("error adding interface", f.Name, e)
				break
			} else if newType != nil {
				ret = append(ret, *newType)
			}
		}
	}
	return
}
