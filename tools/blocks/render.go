package blocks

import (
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	"io"
	"strconv"
)

type Separator func(b *Block, i int) (string, error)

func SpaceSep(b *Block, i int) (ret string, err error) {
	if len(b.Spans) > 0 && (i+1 != len(b.Spans)) {
		ret = " "
	}
	return
}

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

func (b *Block) Render(str io.Writer) (err error) {
	for i, n := range b.Spans {
		if n.Data != nil {
			if s, e := Format(n.Data); e != nil {
				err = e
				break
			} else if _, e := str.Write([]byte(s)); e != nil {
				err = e
				break
			}
			if s, e := n.Sep(b, i); e != nil {
				err = e
				break
			} else if _, e := str.Write([]byte(s)); e != nil {
				err = e
				break
			}
		}
		if n.Block != nil {
			if e := n.Block.Render(str); e != nil {
				err = e
				break
			}
		}
	}
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
		err = errutil.New("Unknown block data type", sbuf.Type{val}, val)
	}
	return
}
