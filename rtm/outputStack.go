package rtm

import (
	"github.com/ionous/sashimi/util/errutil"
	"io"
)

type OutputStack struct {
	output []*PrintMachine
}

func (os *OutputStack) Print(args ...interface{}) (err error) {
	// get the top output, the one we want to write to
	if cnt := len(os.output); cnt > 0 {
		out := os.output[len(os.output)-1]
		err = out.Print(args...)
	} else {
		err = errutil.New("output stack empty")
	}
	return
}

func (os *OutputStack) Println(args ...interface{}) (err error) {
	out := os.output[len(os.output)-1]
	return out.Println(args...)
}

func (os *OutputStack) PushOutput(out io.Writer) {
	os.output = append(os.output, &PrintMachine{flush: out})
}

func (os *OutputStack) PopOutput() {
	os.output = os.output[:len(os.output)-1]
}

func (os *OutputStack) Flush() error {
	out := os.output[len(os.output)-1]
	return out.Flush()
}
