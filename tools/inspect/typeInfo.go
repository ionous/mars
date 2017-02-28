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

func (cmd *CommandInfo) FindParam(name string) (ret ParamInfo, okay bool) {
	for _, p := range cmd.Parameters {
		if p.Name == name {
			ret, okay = p, true
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

func (p *ParamInfo) Usage(parse bool) (uses string, attr map[string]string) {
	parts := strings.Split(p.Uses, "?")
	if len(parts) > 0 {
		uses = parts[0]
		if parse && len(parts) > 1 {
			attr = make(map[string]string)
			for _, q := range strings.Split(parts[1], "&") {
				vs := strings.Split(q, "=")
				attr[vs[0]] = vs[1]
			}
		}
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

// switch to categorization types / constants
// change tests - and fix whatever the f bugs there are.

func (p *ParamInfo) Categorize() (ret ParamType) {
	uses, attr := p.Usage(true)
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
