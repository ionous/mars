package blocks

import (
	"bytes"
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	// "github.com/ionous/mars/export"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	// "github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/recode"
	"github.com/stretchr/testify/assert"
	"strings"
	// "io"
	"testing"
	// "text/tabwriter"
)

func TestEmptyRen(t *testing.T) {
	assert := assert.New(t)
	if types, e := inspect.NewTypes(std.Std()); assert.NoError(e) {
		noun := The("")
		if cmd, ok := types.TypeOf(noun); assert.True(ok) {
			cursor, path := NewDocument(), inspect.NewPath("root")
			if r, e := NewDocBuilder(cursor, path, cmd); assert.NoError(e) {
				if e := inspect.Inspect(types).VisitPath(path, r, noun); assert.NoError(e) {
					doc := cursor.Document()
					assert.Len(doc, 3, "root and 2 children")
				}
			}
		}
		{
			noun := The("stuff", Exists())
			if cmd, ok := types.TypeOf(noun); assert.True(ok) {
				cursor, path := NewDocument(), inspect.NewPath("root")

				// FIX? somethign seems wrnong when weneed to pass the path and noun multiple times.
				if r, e := NewDocBuilder(cursor, path, cmd); assert.NoError(e) {
					if e := inspect.Inspect(types).VisitPath(path, r, noun); assert.NoError(e) {
						doc := cursor.Document()
						assert.Len(doc, 4, "root and 3 children")
					}
				}
			}
		}
	}
}

type WatchBytes struct {
	t       *testing.T
	buf     bytes.Buffer
	lines   []string
	pending string
}

func (w *WatchBytes) String() string {
	return w.buf.String()
}

func (w *WatchBytes) Lines() []string {
	l := w.lines
	if w.pending != "" {
		l = append(l, w.pending)
	}
	return l
}

func (w *WatchBytes) Write(p []byte) (int, error) {
	s := string(p)
	s = strings.Trim(s, "-")
	if len(s) > 0 {
		if s == NewLineString {
			s = "NewLine!"
			w.lines, w.pending = append(w.lines, w.pending), ""
		} else if s == "|" {
			s = strings.Repeat("Tab!", len(s))
		} else {
			w.pending += s
		}
		w.t.Log("wrote:", "'"+s+"'")
	}
	return w.buf.Write(p)
}

func RunManualRules(t *testing.T, what interface{}, rules Rules) (ret string, err error) {
	if types, e := inspect.NewTypes(std.Std()); e != nil {
		err = e
	} else {
		doc := NewDocument()
		if e := BuildDoc(doc, types, what); e != nil {
			err = e
		} else {
			buf := WatchBytes{t: t}
			r := NewRenderer(&buf)
			if e := r.Render(doc.Root(), WatchTerms{t, rules}); e != nil {
				err = e
			} else {
				ret = buf.String()
			}
		}
	}
	return
}

func RunStoryRules(t *testing.T, what interface{}, pack ...*mars.Package) (ret string, err error) {
	pack = append([]*mars.Package{std.Std()}, pack...)
	if types, e := inspect.NewTypes(pack...); e != nil {
		err = e
	} else {
		doc := NewDocument()
		if e := BuildDoc(doc, types, what); e != nil {
			err = e
		} else {
			buf := WatchBytes{t: t}
			r, en := NewRenderer(&buf), NewStoryRules(types)
			if e := r.Render(doc.Root(), WatchTerms{t, en}); e != nil {
				err = e
			} else {
				ret = buf.String()
			}
			// print after to get final ruless
			// text, _ := recode.JsonMarshal(en)
			// t.Log(text)
			// text, _ := recode.JsonMarshal(doc.Root())
			// t.Log(text)
		}
	}
	return
}

func FormatString(data interface{}) (ret string) {
	if data == nil {
		ret = "<blank>"
	} else if s, ok := data.(string); ok {
		ret = s
	} else {
		ret = "<NaS>"
	}
	return
}

type WatchTerms struct {
	t   *testing.T
	src GenerateTerms
}

func (w WatchTerms) GenerateTerms(n *DocNode) TermSet {
	ts := w.src.GenerateTerms(n)
	w.t.Log(n.Path, "generated", len(ts), "terms:")
	for k, old := range ts {
		k, old := k, old // pin these to generate unique variables
		fn := func(data interface{}) string {
			res := old.Filter(data)
			log := res
			if log == NewLineString {
				log = "NewLine!"
			}
			w.t.Log(" ", "`"+log+"`", "from", k.String(), old.Src.String())
			return res
		}
		ts[k] = TermResult{old.Src, fn}
	}
	return ts
}

func TestManualSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	// Prepend, create a rule which produces prefix text for the target.
	Prepend := func(target, text string) *Rule {
		return TermTextWhen(PreTerm, text, IsTarget(target))
	}
	TypeRule := func(name string, fn TermFilter) *Rule {
		return &Rule{"manual type",
			IsParamType{name},
			TermSet{ContentTerm: TermFunction(fn)},
		}
	}

	rules := Rules{
		// thinking a scoped "maker" for things?
		TypeRule("string", FormatString),
		Prepend("NounDirective", "The"),
		Token("man", "NounDirective.Target", "[subject]"),
		Token("man", "NounDirective.Fragments", "[phrases]"),
		// Append("NounDirective", "."),
		TermTextWhen(SepTerm, ".", IsTarget("NounDirective")),
	}

	text, _ := recode.JsonMarshal(rules)
	t.Log(text)

	if text, e := RunManualRules(t, what, rules); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

// generates a noun directive
// directives are used to start describing scripts.
func TestSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	if text, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

func TestExists(t *testing.T) {
	what := The("cabinet", Exists())
	assert := assert.New(t)
	if text, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet exists.", text)
	}
}

func TestKnownAs(t *testing.T) {
	what := The("cabinet", IsKnownAs("the armoire"))
	assert := assert.New(t)
	if text, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet is known as the armoire.", text)
	}
}

func TestUnderstanding(t *testing.T) {
	what := Understand("feed {{something}}").As("feeding it")
	assert := assert.New(t)
	if text, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal(`Understand feed {{something}} as feeding it.`, text)
	}
}

// generates a ScriptRef statement
// ( statements are used in callbacks. )
func TestScriptRef(t *testing.T) {
	what := g.The("fish")
	assert := assert.New(t)
	if text, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("our fish", text)
	}
}

// FIX: evetually these snippets should become part of their test suite
// and we run the matcher externally, generically.
// because ideally, our tests would be near to where they are declared.
func TestIs(t *testing.T) {
	what := g.The("fish").Is("hungry")
	assert := assert.New(t)
	if text, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("is our fish hungry", text)
	}
}

func TestJoinAll(t *testing.T) {
	assert := assert.New(t)
	what := core.All(g.The("fish").Is("hungry"), g.The("fish food").Is("found"))
	if text, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		assert.Equal("( is our fish hungry, and is our fish food found )", text)
	}
}

// FIX? might want to check called in other spots, add rules for prepending "is called" in those cases, and keeping the comma sep.
func TestStoryHeader(t *testing.T) {
	assert := assert.New(t)
	what := The("story",
		Called("A Day for Fresh Sushi"),
		HasText("author", core.T("Emily Short")),
		HasText("headline", core.T("Your basic surreal gay fish romance")),
		Is("scored"),
	)
	if text, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		assert.Equal(`The story called "A Day for Fresh Sushi" has author Emily Short, has headline "Your basic surreal gay fish romance", and is scored.`, text)
	}
}

func xTestSimpleBlock(t *testing.T) {
	assert := assert.New(t)
	what := The("player",
		When("jumping").Always(
			g.Say(`"Er," says the fish. "Does that, like, EVER help??"`),
			g.StopHere(),
		),
		HasText("what", core.T("testing")),
		HasText("what", core.T("booping")),
	)
	if text, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		t.Fatal("\n" + text)
		// assert.Equal(" ", text)
	}
}
