package blocks

import (
	"github.com/ionous/sashimi/util/errutil"
	"io"
	"strconv"
)

// depth can be -1
func (b *Block) RenderToString(maxLen int) (ret string, err error) {
	buf := NewSaturatedOutput(maxLen)
	if e := b.Render(buf); e != nil && e != Saturated {
		err = e
	} else {
		ret = buf.String()
	}
	return
}

type RenderContext struct {
	io.Writer
	space bool
	state State
}

type State struct {
	pos, end int
}

func (rc *RenderContext) Flush(prn bool) {
	if rc.space && prn {
		rc.Write([]byte(" "))
	}
	rc.space = false
}

type ContextRenderer interface {
	ContextRender(*RenderContext) error
	Destroy()
}

func (b *Block) Render(str io.Writer) (err error) {
	return b.ContextRender(&RenderContext{Writer: str})
}

func (b *Block) ContextRender(rc *RenderContext) (err error) {
	memo := rc.state
	rc.state = State{0, len(b.Children)}
	for i, n := range b.Children {
		rc.state.pos = i
		if e := n.ContextRender(rc); e != nil {
			err = e
			break
		}
	}
	rc.state = memo
	return
}

func Format(data interface{}) (ret string, err error) {
	// array of these???
	switch val := data.(type) {
	case string:
		ret = val
	case float64:
		ret = strconv.FormatFloat(val, 'g', -1, 64)
	case bool:
		ret = strconv.FormatBool(val)
	default:
		err = errutil.New("couldnt format data", data)
	}
	return
}
