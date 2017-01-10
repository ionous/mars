package encode

import (
	"github.com/ionous/mars"
	"github.com/ionous/sashimi/util/errutil"
	"net/url"
	r "reflect"
	"strings"
)

type TypeEncoder struct {
	types    TypeBlocks
	faces    Interfaces
	gen      TypeExists
	packages PackageMap
}

func (b *TypeEncoder) Build() TypeBlocks {
	return b.types
}

func NewTypeEncoder() *TypeEncoder {
	return &TypeEncoder{gen: make(TypeExists), packages: make(PackageMap)}
}

type TypeParameters struct {
	Name   string  `json:"name"`
	Phrase *string `json:"phrase,omitempty"`
	Uses   string  `json:"uses"`
}

type TypeBlocks []TypeBlock
type TypeBlock struct {
	Name       string           `json:"name"`
	Implements *string          `json:"implements,omitempty"`
	Parameters []TypeParameters `json:"params,omitempty"`
}

type TypeExists map[r.Type]bool

func packageName(p *mars.Package) (ret *string) {
	if p != nil {
		ret = newString(p.Name)
	}
	return
}
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

func (b *TypeEncoder) addParams(p *mars.Package, s r.Type, ps *Parameters) (err error) {
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

func (b *TypeEncoder) addParam(p *mars.Package, s r.Type, kinds url.Values) (uses string, err error) {
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
		} else if b.faces.Contains(s) {
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
func (b *TypeEncoder) addStruct(p *mars.Package, s r.Type) (err error) {
	if !b.gen[s] {
		if s.Kind() != r.Struct {
			err = errutil.New("not a struct type", s)
		} else {
			b.gen[s] = true
			if face, e := b.faces.FindMatching(s); e != nil {
				err = e
			} else {
				ps := Parameters{}
				if e := b.addParams(p, s, &ps); e != nil {
					err = e
				} else {
					tb := TypeBlock{
						Name:       s.Name(),
						Package:    packageName(p),
						Implements: newString(face),
						Parameters: ps.ps,
					}
					b.types = append(b.types, tb)
				}
			}
		}
	}
	return
}

func (b *TypeEncoder) addInterface(p *mars.Package, t r.Type) (err error) {
	if !b.gen[t] {
		b.gen[t] = true
		name := t.Name()
		tb := TypeBlock{
			Name:       name,
			Package:    packageName(p),
			Implements: newString("interface")}
		b.types = append(b.types, tb)
		b.faces = append(b.faces, InterfaceRecord{name, t})
	}
	return
}

func (b *TypeEncoder) AddPackage(p *mars.Package) (err error) {
	if e := b.addPackage(p); e != nil {
		err = errutil.New("couldn't add package", p.Name, e)
	}
	return
}

func (b *TypeEncoder) addPackage(p *mars.Package) (err error) {
	// l contains all new dependent packages
	if l, e := b.packages.AddPackage(p); e != nil {
		err = e
	} else {
		for _, dep := range l {
			if e := b.addInterfaces(dep, dep.Interfaces); e != nil {
				err = e
				break
			} else if e := b.addCommands(dep, dep.Commands); e != nil {
				err = e
				break
			}
		}
	}
	return
}

func (b *TypeEncoder) addCommands(p *mars.Package, cmds interface{}) (err error) {
	if cmds != nil {
		ref := r.TypeOf(cmds).Elem()
		for i, fields := 0, ref.NumField(); i < fields; i++ {
			f := ref.Field(i)
			elem := f.Type.Elem()
			if e := b.addStruct(p, elem); e != nil {
				err = errutil.New("error adding command", f.Name, e)
				break
			}
		}
	}
	return
}

func (b *TypeEncoder) addInterfaces(p *mars.Package, faces interface{}) (err error) {
	if faces != nil {
		ref := r.TypeOf(faces).Elem()
		for i, fields := 0, ref.NumField(); i < fields; i++ {
			f := ref.Field(i)
			if e := b.addInterface(p, f.Type); e != nil {
				err = errutil.New("error adding interface", f.Name, e)
				break
			}
		}
	}
	return
}
