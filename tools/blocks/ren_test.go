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
	// "strings"
	"testing"
)

func TestEmptyRen(t *testing.T) {
	assert := assert.New(t)
	if types, e := inspect.NewTypes(std.Std()); assert.NoError(e) {
		noun := The("")
		if cmd, ok := types.TypeOf(noun); assert.True(ok) {
			cursor, path := NewDocument(), inspect.NewPath("root")
			if r, e := NewRenderer(cursor, path, cmd); assert.NoError(e) {
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
				if r, e := NewRenderer(cursor, path, cmd); assert.NoError(e) {
					if e := inspect.Inspect(types).VisitPath(path, r, noun); assert.NoError(e) {
						doc := cursor.Document()
						assert.Len(doc, 4, "root and 3 children")
					}
				}
			}
		}
	}
}

func RunManualRules(what interface{}, rules Rules) (ret string, err error) {
	if types, e := inspect.NewTypes(std.Std()); e != nil {
		err = e
	} else {
		doc := NewDocument()
		if e := BuildDoc(doc, types, what); e != nil {
			err = e
		} else {
			var buf bytes.Buffer
			words := NewWordWriter(&buf)
			if e := Render(words, doc.Root(), rules); e != nil {
				err = e
			} else {
				ret = buf.String()
			}
		}
	}
	return
}

type WordWatcher struct {
	t     *testing.T
	words *WordWriter
}

func (ww WordWatcher) WriteWord(s string) {
	ww.t.Logf("write word: '%s'", s)
	ww.words.WriteWord(s)
}

func RunEnglishRules(t *testing.T, what interface{}, pack ...*mars.Package) (ret string, err error) {
	pack = append([]*mars.Package{std.Std()}, pack...)
	if types, e := inspect.NewTypes(pack...); e != nil {
		err = e
	} else {
		doc := NewDocument()
		if e := BuildDoc(doc, types, what); e != nil {
			err = e
		} else {
			var buf bytes.Buffer
			words, en := NewWordWriter(&buf), NewEnglishRules(types)
			if e := Render(words, doc.Root(), en); e != nil {
				err = e
			} else {
				ret = buf.String()
			}
			text, _ := recode.JsonMarshal(en)
			t.Log(text)
			// text, _ = recode.JsonMarshal(doc.Root())
			// t.Log(text)
		}
	}
	return
}

func TestManualSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	rules := Rules{
		// thinking a scoped "maker" for things?
		FormatType("string", FormatString),
		Prepend("NounDirective", "The"),
		Token("NounDirective.Target", "[subject]"),
		Token("NounDirective.Fragments", "[phrases]"),
		Append("NounDirective", "."),
	}
	if text, e := RunManualRules(what, rules); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

// generates a noun directive
// directives are used to start describing scripts.
func TestSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	if text, e := RunEnglishRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

// generates a ScriptRef statement
// ( statements are used in callbacks. )
func TestScriptRef(t *testing.T) {
	what := g.The("fish") //.Is("hungry")
	assert := assert.New(t)
	if text, e := RunEnglishRules(t, what); assert.NoError(e) {
		assert.Equal("our fish", text)
	}
}

func TestExists(t *testing.T) {
	what := The("cabinet", Exists())
	assert := assert.New(t)
	if text, e := RunEnglishRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet exists.", text)
	}
}

func TestKnownAs(t *testing.T) {
	what := The("cabinet", IsKnownAs("the armoire"))
	assert := assert.New(t)
	if text, e := RunEnglishRules(t, what); assert.NoError(e) {
		assert.Equal("The cabinet is known as the armoire.", text)
	}
}

func TestUnderstanding(t *testing.T) {
	what := Understand("feed {{something}}").As("feeding it")
	assert := assert.New(t)
	if text, e := RunEnglishRules(t, what); assert.NoError(e) {
		assert.Equal(`Understand feed {{something}} as feeding it.`, text)
	}
}

// FIX: evetually these snippets should become part of their test suite
// and we run the matcher externally, generically.
// because ideally, our tests would be near to where they are declared.
func TestIs(t *testing.T) {
	what := g.The("fish").Is("hungry")
	assert := assert.New(t)
	if text, e := RunEnglishRules(t, what); assert.NoError(e) {
		assert.Equal("is our fish hungry", text)
	}
}

func TestJoinAll(t *testing.T) {
	assert := assert.New(t)
	what := core.All(g.The("fish").Is("hungry"), g.The("fish food").Is("found"))
	if text, e := RunEnglishRules(t, what, core.Core()); assert.NoError(e) {
		assert.Equal("( is our fish hungry, and is our fish food found )", text)
	}
}
