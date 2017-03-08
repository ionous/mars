package blocks

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/recode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocBuilder(t *testing.T) {
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

func RunManualRules(t *testing.T, what interface{}, rules Rules) (ret string, err error) {
	if types, e := inspect.NewTypes(std.Std()); e != nil {
		err = e
	} else {
		doc := NewDocument()
		if e := BuildDoc(doc, types, what); e != nil {
			err = e
		} else {
			buf := WatchWriter{t: t}
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

func RunStoryRules(t *testing.T, what interface{}, pack ...*mars.Package) (ret *WatchWriter, err error) {
	pack = append([]*mars.Package{std.Std()}, pack...)
	if types, e := inspect.NewTypes(pack...); e != nil {
		err = e
	} else {
		doc := NewDocument()
		if e := BuildDoc(doc, types, what); e != nil {
			err = e
		} else {
			rules := NewStoryRules(types)
			w := WatchWriter{t: t, doc: doc.Document()}
			r := NewRenderer(&w)
			if e := r.Render(doc.Root(), WatchTerms{t, rules}); e != nil {
				err = e
			} else {
				ret = &w
			}
		}
	}
	return
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
	FormatString := func(data interface{}) (ret string) {
		if data == nil {
			ret = "<blank>"
		} else if s, ok := data.(string); ok {
			ret = s
		} else {
			ret = "<NaS>"
		}
		return
	}

	rules := Rules{
		// thinking a scoped "maker" for things?
		TypeRule("string", FormatString),
		Prepend("NounDirective", "The"),
		NewTokenRule("man", "NounDirective.Target", "[subject]"),
		NewTokenRule("man", "NounDirective.Fragments", "[phrases]"),
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
	if out, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", out.String())
	}
}

func TestExists(t *testing.T) {
	what := The("cabinet", Exists())
	assert := assert.New(t)
	if out, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet exists.", out.String())
	}
}

func TestKnownAs(t *testing.T) {
	what := The("cabinet", IsKnownAs("the armoire"))
	assert := assert.New(t)
	if out, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet is known as the armoire.", out.String())
	}
}

func TestUnderstanding(t *testing.T) {
	what := Understand("feed {{something}}").As("feeding it")
	assert := assert.New(t)
	if out, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal(`Understand feed {{something}} as feeding it.`, out.String())
	}
}

// generates a ScriptRef statement
// ( statements are used in callbacks. )
func TestScriptRef(t *testing.T) {
	what := g.The("fish")
	assert := assert.New(t)
	if out, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("our fish", out.String())
	}
}

// FIX: evetually these snippets should become part of their test suite
// and we run the matcher externally, generically.
// because ideally, our tests would be near to where they are declared.
func TestIs(t *testing.T) {
	what := g.The("fish").Is("hungry")
	assert := assert.New(t)
	if out, e := RunStoryRules(t, what); assert.NoError(e) {
		assert.Equal("is our fish hungry", out.String())
	}
}

func TestJoinAll(t *testing.T) {
	assert := assert.New(t)
	what := core.All(g.The("fish").Is("hungry"), g.The("fish food").Is("found"))
	if out, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		assert.Equal("( is our fish hungry, and is our fish food found )", out.String())
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
	if out, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		assert.Equal(`The story called "A Day for Fresh Sushi" has author Emily Short, has headline "Your basic surreal gay fish romance", and is scored.`, out.String())
	}
}

func TestSimpleBlock(t *testing.T) {
	assert := assert.New(t)
	what := The("player",
		When("jumping").Or("clapping").Always(
			g.Say(`"Er," says the fish. "Does that, like, EVER help??"`),
			g.StopHere(),
		),
	)
	if out, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		t.Log(out.String())
		assert.Equal(Lines(
			`The player when jumping or clapping always:`,
			`-Say ""Er," says the fish. "Does that, like, EVER help??""`,
			`-Stop now.`), out.Lines())
	}
}

func TestEmptyBlock(t *testing.T) {
	assert := assert.New(t)
	what := The("player", When("").Always(nil))
	if out, e := RunStoryRules(t, what, core.Core()); assert.NoError(e) {
		// data := out.doc["root/Fragments/0"]
		// t.Log(recode.JsonMarshal(data))
		assert.Equal(Lines(
			"The player when [event name(s)] always:",
			"-[run actions]."), out.Lines())
	}
}
