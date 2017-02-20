package blocks

import (
//	"strings"
)

type Block struct {
	note    BlockNotify
	Path    string
	Kind    string
	Spans   []*Span
	rebuild Rebuild
}
type Rebuild func(*Location) error

// FIX? path should probably be "kid name"
func (b *Block) NewSpan(kind string, path string) *Span {
	id := uniquer
	uniquer++
	n := &Span{Kind: kind, Path: path, Id: id, Sep: SpaceSep}
	b.Spans = append(b.Spans, n)
	return n
}

func (b *Block) Rebuild(loc *Location) error {
	b.destroy(true)
	return b.rebuild(loc)
}

func (b *Block) Destroy() {
	b.destroy(false)
}

func (b *Block) destroy(keep bool) {
	b.Spans = []*Span{}
	if !keep {
		b.note(b)
	}
}

var uniquer int = 0

// func
type Span struct {
	Kind  string
	Data  interface{}
	Path  string
	Id    int
	Sep   Separator
	Block *Block // Spans can hold blocks
}

// func (s *Span) Destroy() {
// }

type Blocks struct {
	script ScriptDB
	blocks map[string]*Block
}

func NewBlocks(s ScriptDB) *Blocks {
	return &Blocks{s, nil}
}

func (l *Blocks) NewRootBlock(path string, kind string, rebuild Rebuild) *Block {
	l.blocks = make(map[string]*Block)
	return l.newBlock(path, kind, rebuild)
}

// FIX: shouldnt path be "kid"
func (l *Blocks) NewChildBlock(parent *Block, path string, kind string, rebuild Rebuild) *Block {
	// create a new child span in the parent
	n := parent.NewSpan(kind, path)
	// create a new block for the passed path
	next := l.newBlock(path, kind, rebuild)
	// put the new block in the new child
	n.Block = next
	// return the new block
	return next
}

func (l *Blocks) newBlock(path string, kind string, rebuild Rebuild) *Block {
	b := &Block{Path: path, Kind: kind, rebuild: rebuild, note: l.blockDestroyed}
	l.blocks[path] = b
	return b
}

func (l *Blocks) Rebuild(path string) (err error) {
	if c, e := l.script.ReverseCursor(path); e != nil {
		err = e
	} else {
		for c.HasNext() {
			loc := c.GetNext()
			if b, ok := l.blocks[loc.Path]; ok {
				if e := b.Rebuild(&loc); e != nil {
					err = e
				}
			}
		}
	}
	return err
}

type BlockNotify func(b *Block)

func (l *Blocks) blockDestroyed(b *Block) {
	delete(l.blocks, b.Path)
}
