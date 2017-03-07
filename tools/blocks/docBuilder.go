package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
)

func BuildDoc(doc *Cursor, types inspect.Types, data interface{}) (err error) {
	if cmd, ok := types.TypeOf(data); !ok {
		err = errutil.New("type not found", data)
	} else {
		path := inspect.NewPath("root")
		if r, e := NewDocBuilder(doc, path, cmd); e != nil {
			err = e
		} else if e := inspect.Inspect(types).VisitPath(path, r, data); e != nil {
			err = e
		} else {
			// the visitor leaves us at the innermost last child,
			// we need to finish all terminal edges.
			err = doc.Flush()
		}
	}
	return
}

func NewDocBuilder(curse *Cursor, path inspect.Path, c *inspect.CommandInfo) (ret *ArgumentsReceiver, err error) {
	next := NewCommandNode(path, c, c, nil)
	if e := curse.Push(next); e != nil {
		err = e
	} else if next.MaxChildren() == 0 {
		_, err = curse.Pop()
	} else {
		ret = &ArgumentsReceiver{c, DocBuilder{curse, nil}}
	}
	return
}

type DocBuilder struct {
	curse *Cursor
	fini  func() error
}

func (l *DocBuilder) dec() (err error) {
	if n := l.curse.Top(); n == nil {
		err = errutil.New("curse empty")
	} else {
		pos, cnt := n.NumChildren(), n.MaxChildren()
		if pos == cnt {
			if _, e := l.curse.Pop(); e != nil {
				err = e
			} else if l.fini != nil {
				err = l.fini()
			}
		}
	}
	return
}

type ArgumentsReceiver struct {
	c *inspect.CommandInfo
	DocBuilder
}
type ElementsReceiver struct {
	b *inspect.CommandInfo
	DocBuilder
}

func NewCommandNode(path inspect.Path, b, c *inspect.CommandInfo, p *inspect.ParamInfo) *DocNode {
	cnt := len(c.Parameters)
	return &DocNode{Path: path, Slot: b, Command: c, Param: p, Children: make([]*DocNode, 0, cnt)}
}

func NewArrayNode(path inspect.Path, b *inspect.CommandInfo, p *inspect.ParamInfo, cnt int) *DocNode {
	return &DocNode{Path: path, Slot: b, Param: p, Children: make([]*DocNode, 0, cnt)}
}

func NewValueNode(path inspect.Path, p *inspect.ParamInfo, d interface{}) *DocNode {
	return &DocNode{Path: path, Param: p, Data: d}
}

func (l *ArgumentsReceiver) NewCommand(path inspect.Path, b, c *inspect.CommandInfo) (ret inspect.Arguments, err error) {
	if p, e := l.findParam(path); e != nil {
		err = e
	} else {
		next := NewCommandNode(path, b, c, p)
		if e := l.curse.Push(next); e != nil {
			err = e
		} else if next.MaxChildren() == 0 {
			err = l.dec()
		} else {
			// after arguments is done, we are done with this command.
			ret = &ArgumentsReceiver{c, DocBuilder{l.curse, func() error {
				return l.dec()
			}}}
		}
	}
	return
}

func (l *ArgumentsReceiver) NewArray(path inspect.Path, b *inspect.CommandInfo, cnt int) (ret inspect.Elements, err error) {
	if p, e := l.findParam(path); e != nil {
		err = e
	} else {
		next := NewArrayNode(path, b, p, cnt)
		if e := l.curse.Push(next); e != nil {
			err = e
		} else if next.MaxChildren() == 0 {
			err = l.dec()
		} else {
			// when elements is done, we can finish
			ret = &ElementsReceiver{b, DocBuilder{l.curse, func() error {
				return l.dec()
			}}}
		}
	}
	return
}

func (l *ArgumentsReceiver) findParam(path inspect.Path) (ret *inspect.ParamInfo, err error) {
	if p, ok := l.c.FindParam(path.Last()); !ok {
		err = errutil.New("couldn't find parameter in", path)
	} else {
		ret = p
	}
	return
}

func (l *ArgumentsReceiver) NewValue(path inspect.Path, d interface{}) (err error) {
	if p, e := l.findParam(path); e != nil {
		err = e
	} else {
		next := NewValueNode(path, p, d)
		if e := l.curse.Push(next); e != nil {
			err = e
		} else {
			err = l.dec()
		}
	}
	return
}

// ElementReceiver NewCommand, adds an element to the array.
func (l *ElementsReceiver) NewElement(path inspect.Path, c *inspect.CommandInfo) (ret inspect.Arguments, err error) {
	next := NewCommandNode(path, l.b, c, nil)
	if e := l.curse.Push(next); e != nil {
		err = e
	} else if next.MaxChildren() == 0 {
		err = l.dec()
	} else {
		ret = &ArgumentsReceiver{c, DocBuilder{l.curse, func() error {
			return l.dec()
		}}}
	}
	return ret, nil
}
