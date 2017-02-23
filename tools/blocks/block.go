package blocks

type Block struct {
	blocks   *Blocks
	Path     string
	Tag      *string           `json:",omitempty"`
	Children []ContextRenderer `json:",omitempty"`
	rebuild  BuildFn
}

func (b *Block) AddSpan(path, tag string) *Span {
	span := NewSpan(path, tag)
	b.Children = append(b.Children, span)
	return span
}

func (b *Block) Build(bk *Stack) error {
	b.destroy(true)
	return b.rebuild(bk)
}

func (b *Block) Destroy() {
	b.destroy(false)
}

func (b *Block) destroy(keep bool) {
	for _, n := range b.Children {
		n.Destroy()
	}
	b.Children = nil
	if !keep {
		b.blocks.blockDestroyed(b)
	}
}
