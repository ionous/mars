package blocks

// Generator: implements DocStack;
// runs rules immediately upon push/pop
type Generator struct {
	Writer Words
	Rules  RuleFinder
	*DocumentCursor
}

func (g *Generator) Push(n *DocNode) (err error) {
	if e := g.DocumentCursor.Push(n); e != nil {
		err = e
	} else if e := g.runRules(n, ApplyBefore); e != nil {
		err = e
	} else {
		err = g.runRules(n, ApplyOn)
	}
	return
}

func (g *Generator) Pop() (ret *DocNode, err error) {
	if n, e := g.DocumentCursor.Pop(); e != nil {
		err = e
	} else if e := g.runRules(n, ApplyAfter); e != nil {
		err = e
	}
	return
}

func (g *Generator) runRules(n *DocNode, when ApplyWhen) (err error) {
	if r, ok := g.Rules.FindBestRule(MatchSource{n, when}); ok {
		if e := r.Write(g.Writer, n); e != nil {
			err = e
		}
	}
	return
}
