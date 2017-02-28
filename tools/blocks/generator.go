package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
)

func Render(words Words, node *DocNode, rules RuleFinder) (err error) {
	for _, n := range node.Children {
		if e := runRules(words, n, rules, ApplyBefore); e != nil {
			err = e
			break
		} else if e := runRules(words, n, rules, ApplyOn); e != nil {
			err = e
			break
		} else if e := Render(words, n, rules); e != nil {
			err = e
			break
		} else if e := runRules(words, n, rules, ApplyAfter); e != nil {
			break
		}
	}
	return
}

func runRules(words Words, n *DocNode, rules RuleFinder, when ApplyWhen) (err error) {
	if r, ok := rules.FindBestRule(MatchSource{n, when}); ok {
		if e := r.Write(words, n); e != nil {
			err = e
		}
	}
	return
}

func BuildDoc(doc DocStack, types inspect.Types, data interface{}) (err error) {
	if cmd, ok := types.TypeOf(data); !ok {
		err = errutil.New("type not found", data)
	} else {
		var path inspect.Path
		if r, e := NewRenderer(doc, path, cmd); e != nil {
			err = e
		} else if e := inspect.Inspect(types).VisitPath(path, r, data); e != nil {
			err = e
		} else {
			// the visitor leaves us at the innermost last child,
			// we need to finish all terminal edges.
			err = PopStack(doc)
		}
	}
	return
}
