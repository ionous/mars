package encode

import (
	"github.com/ionous/mars"
	"github.com/ionous/sashimi/util/errutil"
	"net/url"
	r "reflect"
	"strings"
)

type TypeRecoder struct {
	allFaces Interfaces
	Types    TypeMap
}

func NewTypeRecoder() *TypeRecoder {
	return &TypeRecoder{
		Types: make(TypeMap),
	}
}

func (b *TypeRecoder) AddTypes(ps ...*mars.Package) (ret []CommandType, err error) {
	for _, p := range ps {
		if faceTypes, e := b.addInterfaces(p, p.Interfaces); e != nil {
			err = e
			break
		} else if cmdTypes, e := b.addCommands(p, p.Commands); e != nil {
			err = e
			break
		} else {
			ret = append(faceTypes, cmdTypes...)
		}
	}
	return
}

type ParamInfo struct {
	Name   string  `json:"name"`
	Phrase *string `json:"phrase,omitempty"`
	Uses   string  `json:"uses"`
}

type CommandType struct {
	Name       string      `json:"name"`
	Implements *string     `json:"implements,omitempty"`
	Parameters []ParamInfo `json:"params,omitempty"`
}

func (cmd *CommandType) FindParam(name string) (ret ParamInfo, okay bool) {
	for _, p := range cmd.Parameters {
		if p.Name == name {
			ret, okay = p, true
			break
		}
	}
	return
}

type TypeMap map[string]*CommandType

func (p *ParamInfo) Split() (uses string, style map[string]string) {
	parts := strings.Split(p.Uses, "?")
	if len(parts) > 0 {
		uses = parts[0]
		if len(parts) > 1 {
			style = make(map[string]string)
			for _, q := range strings.Split(parts[1], "&") {
				vs := strings.Split(q, "=")
				style[vs[0]] = vs[1]
			}
		}
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
	ps []ParamInfo
}

func (b *TypeRecoder) addParams(p *mars.Package, s r.Type, ps *Parameters) (err error) {
	if s.Kind() != r.Struct {
		err = errutil.New("couldn't add params of", s)
	} else {
		for i, cnt := 0, s.NumField(); i < cnt; i++ {
			f := s.Field(i)
			// pkg path is empty only for public members
			if f.PkgPath == "" {
				tp := ParamInfo{}
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
	}
	return
}

func (b *TypeRecoder) addParam(p *mars.Package, s r.Type, kinds url.Values) (uses string, err error) {
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
	case r.Struct:
		uses = "blob"
		kinds.Add("type", n)

	default:
		err = errutil.New("has unsupported", k)
	}
	return
}

// meeds a bit of recursion.
func (b *TypeRecoder) addStruct(p *mars.Package, s r.Type) (ret *CommandType, err error) {
	if s.Kind() != r.Struct {
		err = errutil.New("not a struct type", s)
	} else {
		name := s.Name()
		if b.Types[name] == nil {
			if face, e := b.allFaces.FindMatching(s); e != nil {
				err = e
			} else {
				ps := Parameters{}
				if e := b.addParams(p, s, &ps); e != nil {
					err = e
				} else {
					typeInfo := &CommandType{
						Name:       name,
						Implements: newString(face),
						Parameters: ps.ps,
					}
					b.Types[name] = typeInfo
					ret = typeInfo
				}
			}
		}
	}
	return
}

func (b *TypeRecoder) addInterface(p *mars.Package, t r.Type) (ret *CommandType, err error) {
	name := t.Name()
	if b.Types[name] == nil {
		b.allFaces = append(b.allFaces, InterfaceRecord{name, t})
		typeInfo := &CommandType{
			Name:       name,
			Implements: newString("interface"),
		}
		b.Types[name] = typeInfo
		ret = typeInfo
	}
	return
}

func (b *TypeRecoder) addCommands(p *mars.Package, cmds interface{}) (ret []CommandType, err error) {
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

func (b *TypeRecoder) addInterfaces(p *mars.Package, allFaces interface{}) (ret []CommandType, err error) {
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
