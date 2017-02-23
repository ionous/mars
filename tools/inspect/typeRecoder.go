package inspect

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/tools/inspect/internal"
	"github.com/ionous/sashimi/util/errutil"
	"net/url"
	r "reflect"
	"strings"
)

type TypeRecoder struct {
	allFaces internal.Interfaces
	Types    Types
}

func NewTypeRecoder() *TypeRecoder {
	return &TypeRecoder{
		Types: make(Types),
	}
}

func NewTypes(ps ...*mars.Package) (ret Types, err error) {
	pm := PackageMap{}
	if pl, e := pm.AddPackages(ps...); e != nil {
		err = e
	} else {
		tr := &TypeRecoder{
			Types: make(Types),
		}
		if e := tr.addPackageList(pl); e != nil {
			err = e
		} else {
			ret = tr.Types
		}
	}
	return
}

func (tr *TypeRecoder) AddTypes(p *mars.Package) (ret []CommandInfo, err error) {
	if faceTypes, e := tr.addInterfaces(p, p.Interfaces); e != nil {
		err = e
	} else if cmdTypes, e := tr.addCommands(p, p.Commands); e != nil {
		err = e
	} else {
		ret = append(faceTypes, cmdTypes...)
	}
	return
}

func (tr *TypeRecoder) addPackageList(ps PackageList) (err error) {
	for _, p := range ps {
		if _, e := tr.AddTypes(p); e != nil {
			err = e
			break
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

func ParseMarsTag(f *r.StructField) (phrase, primType string) {
	if tags := f.Tag.Get("mars"); tags != "" {
		phraseType := strings.Split(tags, ";")
		phrase = phraseType[0]
		if cnt := len(phraseType); cnt == 2 {
			primType = phraseType[1]
		}
	}
	return
}

func (tr *TypeRecoder) addParams(p *mars.Package, s r.Type, ps *Parameters) (err error) {
	if s.Kind() != r.Struct {
		err = errutil.New("couldn't add params of", s)
	} else {
		for i, cnt := 0, s.NumField(); i < cnt; i++ {
			f := s.Field(i)
			// pkg path is empty only for public members
			if f.PkgPath == "" {
				tp := ParamInfo{}
				tp.Name = f.Name
				phrase, primType := ParseMarsTag(&f)
				tp.Phrase = newString(phrase)
				//
				if f.Anonymous {
					if e := tr.addParams(p, f.Type, ps); e != nil {
						err = e
						break
					}
				} else {
					kinds := make(url.Values)
					if uses, e := tr.addParam(p, f.Type, kinds); e != nil {
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

func (tr *TypeRecoder) addParam(p *mars.Package, s r.Type, kinds url.Values) (uses string, err error) {
	switch n, k := s.Name(), s.Kind(); k {
	case r.String, r.Bool, r.Float64:
		uses = k.String()
		if n != uses {
			kinds.Add("type", n)
		}

	case r.Array, r.Slice:
		uses, err = tr.addParam(p, s.Elem(), kinds)
		kinds.Add("array", "true")

	case r.Interface:
		// FIX: for now.
		if n == "Generic" {
			uses = "ObjEval"
		} else if tr.allFaces.Contains(s) {
			uses = n
		} else if n != "" {
			err = errutil.New("has unknown interface", n)
		} else {
			uses = "blob"
		}
	case r.Struct:
		uses = "blob"
		kinds.Add("type", n)

	default:
		err = errutil.New("has unsupported", k)
	}
	return
}

// needs a bit of recursion.
func (tr *TypeRecoder) addStruct(p *mars.Package, f *r.StructField) (ret *CommandInfo, err error) {
	if s := f.Type.Elem(); s.Kind() != r.Struct {
		err = errutil.New("not a struct type", s)
	} else {
		name := s.Name()
		if tr.Types[name] == nil {
			if face, e := tr.allFaces.FindMatching(s); e != nil {
				err = e
			} else {
				ps := Parameters{}
				if e := tr.addParams(p, s, &ps); e != nil {
					err = e
				} else {
					phrase, category := ParseMarsTag(f)
					typeInfo := &CommandInfo{
						Name:       name,
						Implements: newString(face),
						Phrase:     newString(phrase),
						Category:   newString(category),
						Parameters: ps.ps,
					}
					tr.Types[name] = typeInfo
					ret = typeInfo
				}
			}
		}
	}
	return
}

func (tr *TypeRecoder) addInterface(p *mars.Package, t r.Type) (ret *CommandInfo, err error) {
	if name := t.Name(); tr.Types[name] == nil {
		tr.allFaces = append(tr.allFaces, internal.NewInterface(t))
		typeInfo := &CommandInfo{
			Name:       name,
			Implements: newString("interface"),
		}
		tr.Types[name] = typeInfo
		ret = typeInfo
	}
	return
}

func (tr *TypeRecoder) addCommands(p *mars.Package, cmds interface{}) (ret []CommandInfo, err error) {
	if cmds != nil {
		ref := r.TypeOf(cmds).Elem()
		for i, fields := 0, ref.NumField(); i < fields; i++ {
			f := ref.Field(i)
			if newType, e := tr.addStruct(p, &f); e != nil {
				err = errutil.New("error adding command", f.Name, e)
				break
			} else if newType != nil {
				ret = append(ret, *newType)
			}
		}
	}
	return
}

func (tr *TypeRecoder) addInterfaces(p *mars.Package, allFaces interface{}) (ret []CommandInfo, err error) {
	if allFaces != nil {
		ref := r.TypeOf(allFaces).Elem()
		for i, fields := 0, ref.NumField(); i < fields; i++ {
			f := ref.Field(i)
			if newType, e := tr.addInterface(p, f.Type); e != nil {
				err = errutil.New("error adding interface", f.Name, e)
				break
			} else if newType != nil {
				ret = append(ret, *newType)
			}
		}
	}
	return
}
