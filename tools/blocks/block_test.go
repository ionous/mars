package blocks

import (
	"bytes"
	"encoding/json"
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/export"
	"github.com/ionous/mars/script"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

var _ = bytes.Equal
var _ = log.Println
var _ = export.Export
var _ = json.Indent

//
func PhraseText(what interface{}, pack ...*mars.Package) (ret string, err error) {
	pack = append([]*mars.Package{std.Std()}, pack...)
	if types, e := inspect.NewTypes(pack...); e != nil {
		err = e
	} else if db, e := NewDBMaker("test", types).Compute(what); e != nil {
		err = e
	} else {
		// prettyBytes, _ := json.MarshalIndent(db, "", " ")
		// log.Println("db", string(prettyBytes))

		m := NewStoryModel(db, types)
		if block, _, e := m.BuildRootCmd("test"); e != nil {
			err = e
		} else {
			prettyBytes, _ := json.MarshalIndent(*block, "", " ")
			log.Println("blocks", string(prettyBytes))

			var buf bytes.Buffer
			if e := block.Render(&buf); e != nil {
				err = e
			} else {
				ret = buf.String()
			}
		}
	}
	return
}

func TestStrings(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("fragment", PascalSpaces("Fragment"))
	assert.Equal("fragment fragment", PascalSpaces("FragmentFragment"))
}

func TestTokenize(t *testing.T) {
	assert := assert.New(t)
	if true {
		pre, post, token := TokenizePhrase("[phrases]")
		actual := strings.Join([]string{pre, post, token}, ";")
		assert.Equal(";;[phrases]", actual)
	}
	if true {
		pre, post, token := TokenizePhrase("The [subject]")
		actual := strings.Join([]string{pre, post, token}, ";")
		assert.Equal("The;;[subject]", actual)
	}
	if true {
		pre, post, token := TokenizePhrase("The [noun] uses")
		actual := strings.Join([]string{pre, post, token}, ";")
		assert.Equal("The;uses;[noun]", actual)
	}
	if true {
		pre, post, token := TokenizePhrase("nope")
		actual := strings.Join([]string{pre, post, token}, ";")
		assert.Equal("nope;;", actual)
	}
}

func TestParamTypes(t *testing.T) {
	assert := assert.New(t)
	if types, e := inspect.NewTypes(std.Std()); assert.NoError(e) {
		if cmd, ok := types["NounDirective"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Target"); assert.True(ok) {
				r := p.Categorize()
				assert.Equal(inspect.ParamTypePrim, r, r.String())
			}
			if p, ok := cmd.FindParam("Fragments"); assert.True(ok) {
				r := p.Categorize()
				assert.Equal(inspect.ParamTypeArray, r, r.String())
			}
		}
		if cmd, ok := types["ScriptRef"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Obj"); assert.True(ok) {
				r := p.Categorize()
				assert.Equal(inspect.ParamTypeCommand, r, r.String())
			}
		}
	}
	if types, e := inspect.NewTypes(export.Export()); assert.NoError(e) {
		if cmd, ok := types["Library"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Types"); assert.True(ok) {
				r := p.Categorize()
				assert.Equal(inspect.ParamTypeBlob, r, r.String())
			}
		}
	}
	if types, e := inspect.NewTypes(script.Package()); assert.NoError(e) {
		if cmd, ok := types["ParserDirective"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Input"); assert.True(ok) {
				r := p.Categorize()
				assert.Equal(inspect.ParamTypeArray, r, r.String())
			}
		}
	}
}

// generates a noun directive
// directives are used to start describing scripts.
func TestSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

// generates a ScriptRef statement
// statements are used in callbacks.
func TestScriptRef(t *testing.T) {
	what := g.The("fish") //.Is("hungry")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("our fish", text)
	}
}

func TestExists(t *testing.T) {
	what := The("cabinet", Exists()) //IsKnownAs("armoire")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("The cabinet exists.", text)
	}
}

func TestKnownAs(t *testing.T) {
	what := The("cabinet", IsKnownAs("the armoire"))
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("The cabinet is known as the armoire.", text)
	}
}

func TestUnderstanding(t *testing.T) {
	what := Understand("feed {{something}}").As("feeding it")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal(`Understand feed {{something}} as feeding it.`, text)
	}
}

// FIX: evetually these snippets should become part of their test suite
// and we run the matcher externally, generically.
// because ideally, our tests would be near to where they are declared.
func TestIs(t *testing.T) {
	what := g.The("fish").Is("hungry")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("is our fish hungry", text)
	}
}

//
func TestJoinAll(t *testing.T) {
	assert := assert.New(t)
	what := core.All(g.The("fish").Is("hungry"), g.The("fish food").Is("found"))
	assert.Len(what.(core.AllTrue).Test, 2)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("( is our fish hungry, and is our fish food found )", text)
	}
}
