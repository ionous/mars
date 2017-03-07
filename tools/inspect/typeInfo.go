package inspect

import (
	r "reflect"
	"strings"
)

type ParamInfo struct {
	Name   string  `json:"name"`
	Phrase *string `json:"phrase,omitempty"`
	Uses   string  `json:"uses"`
}

type CommandInfo struct {
	Name string `json:"name"`
	// FIX: currently an optional, comma separated string
	// change to: an array
	Implements *string     `json:"implements,omitempty"`
	Parameters []ParamInfo `json:"params,omitempty"`
	Phrase     *string     `json:"phrase,omitempty"`
	Category   *string     `json:"category,omitempty"`
}

type Types map[string]*CommandInfo

func (t Types) TypeOf(data interface{}) (ret *CommandInfo, okay bool) {
	if data != nil {
		name := r.TypeOf(data).Name()
		if f, ok := t[name]; ok {
			ret, okay = f, true
		}
	}
	return
}

func (cmd *CommandInfo) FindParam(name string) (ret *ParamInfo, okay bool) {
	for _, p := range cmd.Parameters {
		if p.Name == name {
			p := p // pin
			ret, okay = &p, true
			break
		}
	}
	return
}

func (cmd *CommandInfo) Types() (ret []string) {
	if cmd.Implements == nil {
		ret = []string{cmd.Name}
	} else {
		i := strings.Split(*cmd.Implements, ",")
		ret = append([]string{cmd.Name}, i...)
	}
	return
}

func (p *ParamInfo) Type() string {
	t, _ := p.Usage(false)
	return t
}

type ParamUsage struct {
	parts []string
	attr  map[string]string
}

func (u *ParamUsage) Uses() (ret string) {
	if len(u.parts) > 0 {
		ret = u.parts[0]
	}
	return
}

func (u *ParamUsage) Attrs() (ret map[string]string) {
	if u.attr == nil && len(u.parts) > 1 {
		u.attr = make(map[string]string)
		for _, q := range strings.Split(u.parts[1], "&") {
			vs := strings.Split(q, "=")
			u.attr[vs[0]] = vs[1]
		}
	}
	return u.attr
}

func (p *ParamInfo) ParamUsage() ParamUsage {
	parts := strings.Split(p.Uses, "?")
	return ParamUsage{parts: parts}
}

func (p *ParamInfo) Usage(parse bool) (uses string, attr map[string]string) {
	x := p.ParamUsage()
	uses = x.Uses()
	if parse {
		attr = x.Attrs()
	}
	return
}

//go:generate stringer -type=ParamType
type ParamType int

const (
	ParamTypeUnknown ParamType = iota
	ParamTypeCommand
	ParamTypeArray
	ParamTypePrim
	ParamTypeBlob
)

func (p *ParamInfo) Categorize() ParamType {
	x := p.ParamUsage()
	return x.Category()
}

func (u *ParamUsage) Category() (ret ParamType) {
	uses, attr := u.Uses(), u.Attrs()
	if uses == "blob" {
		ret = ParamTypeBlob
	} else if strings.ToUpper(uses[:1]) != uses[:1] {
		ret = ParamTypePrim
	} else if attr["array"] == "true" {
		ret = ParamTypeArray
	} else {
		ret = ParamTypeCommand
	}
	return
}
