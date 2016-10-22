package rtm

import (
	"bytes"
	"fmt"
	"io"
	"unicode"
)

type PrintMachine struct {
	buf   bytes.Buffer
	flush io.Writer
}

// Say("hello", "there.", "world."),
// becomes "hello there. world.\n"
func (p *PrintMachine) Print(args ...interface{}) (err error) {
	var temp bytes.Buffer
	if n, e := fmt.Fprint(&temp, args...); e != nil {
		err = e
	} else if n > 0 {
		s := temp.String()
		// printed something before?
		if p.buf.Len() != 0 {
			// before writing this new thing, possibly put a space.
			run := []rune(s)[0]
			if !unicode.IsPunct(run) && !unicode.IsSpace(run) {
				p.buf.WriteString(" ")
			}
		}
		_, err = p.buf.WriteString(s)
	}
	return
}

func (p *PrintMachine) Println(args ...interface{}) (err error) {
	if _, e := fmt.Fprintln(&p.buf, args...); e != nil {
		err = e
	} else {
		err = p.Flush()
	}
	return
}

func (p *PrintMachine) Flush() (err error) {
	if p.buf.Len() > 0 {

		if _, e := p.flush.Write(p.buf.Bytes()); e != nil {
			err = e
		} else {
			p.buf.Reset()
		}
	}
	return
}
