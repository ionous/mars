package blocks

import (
	"bytes"
	"github.com/ionous/sashimi/util/errutil"
)

const Saturated = errutil.Error("github.com.ionous.mars.tools.blocks.saturated")

// implements BlockOutput
type Sat struct {
	buf    bytes.Buffer
	maxLen int
}

func NewSaturatedOutput(maxLen int) *Sat {
	return &Sat{maxLen: maxLen}
}

func (o *Sat) String() string {
	return o.buf.String()
}

func (o *Sat) Write(p []byte) (ret int, err error) {
	if o.maxLen < 0 {
		ret, err = o.buf.Write(p)
	} else {
		if rem := o.maxLen - o.buf.Len(); rem > len(p) {
			ret, err = o.buf.Write(p)
		} else {
			ret, err = o.buf.Write(p[:rem])
			err = Saturated
		}
	}
	return
}
