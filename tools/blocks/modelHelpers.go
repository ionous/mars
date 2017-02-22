package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type BlockTags map[string]string
type ParamProcs map[string]BuildFn
type CommandProcs map[string]CommandFn

func (m BlockTags) CreateOn(typeName string, tag string) {
	m[typeName] = tag
}

func (m CommandProcs) AddProcess(uses string, proc CommandFn) (err error) {
	if _, exists := m[uses]; !exists {
		m[uses] = proc
	} else {
		err = errutil.New("duplicate process", uses)
		panic(err)
	}
	return
}

func (m ParamProcs) AddProcess(uses string, proc BuildFn) (err error) {
	if _, exists := m[uses]; !exists {
		m[uses] = proc
	} else {
		err = errutil.New("duplicate process", uses)
		panic(err)
	}
	return
}

func (m ParamProcs) AddFilter(uses string, filter FilterFn) error {
	return m.AddProcess(uses, func(stack *Stack) error {
		return stack.Data(func(data interface{}) (err error) {
			if src, e := Format(data.(*PrimData).Value); e != nil {
				err = errutil.New("error filtering", uses, e)
			} else {
				stack.NewSpan("st-prim", func(span *Span) {
					span.Text = filter(src)
				})
			}
			return
		})
	})
}

func (m *ModelMaker) commandFromData(data interface{}) (ret *inspect.CommandInfo, err error) {
	if src, ok := data.(*CommandData); !ok {
		err = errutil.New("not a command", sbuf.Type{data})
	} else {
		var name string
		if len(src.Cmd) > 0 {
			name = src.Cmd
		} else if len(src.Type) > 0 {
			name = src.Type
		}
		if name == "" {
			err = errutil.New("empty command")
		} else if cmd, ok := m.types[name]; !ok {
			err = errutil.New("unknown command", name)
		} else {
			ret = cmd
		}
	}
	return
}

func (m *ModelMaker) innerBuild(stack *Stack) error {
	return stack.Command(func(cmd *inspect.CommandInfo) (err error) {
		if len(cmd.Parameters) == 0 {
			spaces := PascalSpaces(cmd.Name)
			stack.NewSpan("st-cmd", func(span *Span) {
				span.Text = spaces
			})
		} else {
			err = stack.NewParameters(func(param *inspect.ParamInfo) error {
				path := stack.ChildPath(param.Name)
				return stack.NewPath(path, func(data interface{}) (err error) {
					pre, post, token := Tokenize(param)
					if data == nil {
						stack.NewSpan("st-token", func(span *Span) {
							span.Text = token
						})
					} else {
						if len(pre) > 0 {
							stack.NewSpan("st-pre", func(span *Span) {
								span.Text = pre
							})
						}
						dotted := cmd.Name + "." + param.Name
						var proc BuildFn
						if p, ok := m.Params[dotted]; ok {
							proc = p
						} else if p, ok := m.Params[param.Uses]; ok {
							proc = p
						} else {
							proc = m.buildContent
						}
						if e := proc(stack); e != nil {
							err = e
						} else {
							if len(post) > 0 {
								stack.NewSpan("st-post", func(span *Span) {
									span.Text = post
								})
							}
						}
					}
					return
				})
			})
		}
		return
	})
}

//
func (m *ModelMaker) buildContent(stack *Stack) (err error) {
	return stack.Parameter(func(param *inspect.ParamInfo) (err error) {
		switch param.Categorize() {
		case inspect.ParamTypePrim:
			err = m.BuildPrimitive(stack)
		case inspect.ParamTypeArray:
			err = m.BuildArray(stack)
		case inspect.ParamTypeBlob:
			//
		case inspect.ParamTypeCommand:
			err = m.BuildCmd(stack)
		default:
			err = errutil.New("unknown primitive type")
		}
		if err != nil {
			err = errutil.New("couldnt build content", param.Name, err)
		}
		return
	})
	return
}
