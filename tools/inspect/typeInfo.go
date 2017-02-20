package inspect

import "strings"

type ParamInfo struct {
	Name   string  `json:"name"`
	Phrase *string `json:"phrase,omitempty"`
	Uses   string  `json:"uses"`
}

type CommandInfo struct {
	Name       string      `json:"name"`
	Implements *string     `json:"implements,omitempty"`
	Parameters []ParamInfo `json:"params,omitempty"`
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

type Type map[string]*CommandInfo

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
