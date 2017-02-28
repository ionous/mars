package blocks

type DocStack interface {
	Top() *DocNode
	Push(*DocNode) error
	Pop() (*DocNode, error)
}

// keep popping the stack till empty
func PopStack(g DocStack) (err error) {
	for g.Top() != nil {
		if _, e := g.Pop(); e != nil {
			err = e
			break
		}
	}
	return
}
