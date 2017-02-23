package blocks

import (
	"strings"
	"unicode"
)

type Span struct {
	Tag      string
	Text     string
	Path     string
	Sep      Separator `json:",omitempty"`
	Children *Block    `json:",omitempty"`
}

func NewSpan(path, tag string) *Span {
	// uniquify the path
	if len(tag) > 0 {
		path += "?" + strings.Replace(tag, " ", "&", -1)
	}
	return &Span{
		Path: path,
		Tag:  tag,
	}
}

func (n *Span) Destroy() {
}

func (n *Span) ContextRender(rc *RenderContext) (err error) {
	if s := strings.TrimSpace(n.Text); len(s) > 0 {
		rc.Flush(len(strings.TrimFunc(s, unicode.IsPunct)) > 0)
		if _, e := rc.Write([]byte(s)); e != nil {
			err = e
			return
		}
		rc.space = true
		// soft space:
		if seps := n.Sep; seps == nil {
			rc.space = true
		} else {
			// if there is a sep -- it takes over and replaces space.
			if sep := seps.Sep(&rc.state); len(sep) > 0 {
				// rc.Flush(false)
				if _, e := rc.Write([]byte(sep)); e != nil {
					err = e
					return
				}
			}
		}
	}
	if n.Children != nil {
		if e := n.Children.ContextRender(rc); e != nil {
			err = e
			return
		}
	}
	return
}
