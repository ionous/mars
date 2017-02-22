package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	// "log"
)

type Stack struct {
	db        ScriptDB
	path      string
	data      interface{}
	block     *Block
	blocks    *Blocks
	cmd       *inspect.CommandInfo
	param     *inspect.ParamInfo
	LastChild *Span
}

func NewStack(db ScriptDB, blocks *Blocks) *Stack {
	return &Stack{
		db:     db,
		blocks: blocks,
	}
}

// Path, the current path.
func (bk *Stack) Path() string {
	return bk.path
}

// ChildPath, the current path plus some kid
func (bk *Stack) ChildPath(kid string) string {
	return SlashPath(bk.path, kid)
}

// Block, callback with the current block.
func (bk *Stack) Block(cb func(*Block) error) error {
	return cb(bk.block)
}

// Command, callback with the current command.
func (bk *Stack) Command(cb func(*inspect.CommandInfo) error) error {
	return cb(bk.cmd)
}

// Parameter, callback with the current param, errors if no parameter scope set.
func (bk *Stack) Parameter(cb func(*inspect.ParamInfo) error) (err error) {
	if bk.param == nil {
		err = errutil.New("no parameter scope at", bk.path)
	} else {
		err = cb(bk.param)
	}
	return
}

// Data, callback with the current data; can be null.
func (bk *Stack) Data(cb func(interface{}) error) error {
	return cb(bk.data)
}

// NewPath, establish a new data context.
func (bk *Stack) NewPath(newPath string, cb func(interface{}) error) error {
	var newData interface{}
	if data, ok := bk.db.FindByPath(newPath); ok {
		newData = data
	}
	oldPath, oldData := bk.path, bk.data
	bk.path, bk.data = newPath, newData
	err := cb(bk.data)
	bk.path, bk.data = oldPath, oldData
	return err
}

func (bk *Stack) NewRoot(tag string, cb func(*Stack) error) (ret *Block, err error) {
	if bk.path == "" {
		err = errutil.New("blocks need paths")
	} else if bk.block != nil {
		err = errutil.New("root already created")
	} else {
		block := bk.blocks.newBlock(bk.path, tag, cb)
		bk.block = block
		if e := cb(bk); e != nil {
			err = e
		} else {
			ret = block
		}
		bk.block = nil
	}
	return
}

// NewBlock, create a new block with passed tag.
func (bk *Stack) NewBlock(tag string, cb func(*Stack) error) error {
	// blocks live in spans, so first create one:
	span := bk.block.AddSpan(bk.path, tag)
	// create a new block for the passed path
	newBlock, oldBlock := bk.blocks.newBlock(bk.path, tag, cb), bk.block
	// put the new block in the new child
	span.Children = newBlock
	//
	bk.block = newBlock
	err := cb(bk)
	bk.block = oldBlock
	return err
}

// NewSpan, create a new span with passed tag;
// should always succeed, or panic.
func (bk *Stack) NewSpan(tag string, cb func(*Span)) {
	if bk.block == nil {
		panic("no block")
	}
	span := bk.block.AddSpan(bk.path, tag)
	cb(span)
	bk.LastChild = span
}

// NewCommand, establish a new command scope.
func (bk *Stack) NewCommand(cmd *inspect.CommandInfo, cb func() error) error {
	newCmd, oldCmd := cmd, bk.cmd
	bk.cmd = newCmd
	err := cb()
	bk.cmd = oldCmd
	return err
}

// NewParameters, establish a series of parameter scopes based on the current comand.
func (bk *Stack) NewParameters(cb func(*inspect.ParamInfo) error) (err error) {
	cmd, oldParam := bk.cmd, bk.param
	for _, p := range cmd.Parameters {
		bk.param = &p
		if e := cb(bk.param); e != nil {
			err = e
			break
		}
	}
	bk.param = oldParam
	return
}
