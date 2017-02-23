package blocks

type Blocks struct {
	db     ScriptDB
	blocks map[string]*Block
}

func NewBlocks(db ScriptDB) *Blocks {
	return &Blocks{db, make(map[string]*Block)}
}

func (l *Blocks) newBlock(path, tag string, rebuild BuildFn) *Block {
	var opt *string
	if len(tag) > 0 {
		opt = &tag
	}
	b := &Block{
		blocks:  l,
		Path:    path,
		Tag:     opt,
		rebuild: rebuild,
	}
	l.blocks[path] = b
	return b
}

func (l *Blocks) blockDestroyed(b *Block) {
	delete(l.blocks, b.Path)
}

// func (l *Blocks) Build(bk *Stack, path string) (err error) {
// 	if c, e := l.db.ReverseCursor(path); e != nil {
// 		err = e
// 	} else {
// 		// search upwards for the first path
// 		// then we will build downwards
// 		for c.HasNext() {
// 			loc := c.GetNext()
// 			if b, ok := l.blocks[loc.Path]; ok {
// 				if e := b.BuildFn(&loc); e != nil {
// 					err = e
// 				}
// 				break
// 			}
// 		}
// 	}
// 	return err
// }
