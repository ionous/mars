package blocks

type Block struct {
	blocks  *Blocks
	Path    string
	Tag     string
	Spans   []*Span `json:",omitempty"`
	rebuild BuildFn
}

func (b *Block) AddSpan(path, tag string) *Span {
	n := &Span{
		Path: path + "?" + tag,
		Tag:  tag,
		Sep:  SpaceSep{},
	}
	b.Spans = append(b.Spans, n)
	return n
}

func (b *Block) Build(bk *Stack) error {
	b.destroy(true)
	return b.rebuild(bk)
}

func (b *Block) Destroy() {
	b.destroy(false)
}

func (b *Block) destroy(keep bool) {
	b.Spans = []*Span{}
	if !keep {
		b.blocks.blockDestroyed(b)
	}
}
