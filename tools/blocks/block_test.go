package blocks

import (
	"bytes"
	"encoding/json"
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/export"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var _ = bytes.Equal
var _ = log.Println
var _ = export.Export
var _ = json.Indent

func Marshal(src interface{}) (ret string, err error) {
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")
	if e := enc.Encode(src); e != nil {
		err = e
	} else {
		ret = b.String()
	}
	return
}

func PhraseText(what interface{}, pack ...*mars.Package) (ret string, err error) {
	pack = append([]*mars.Package{std.Std()}, pack...)
	if types, e := inspect.NewTypes(pack...); e != nil {
		err = e
	} else if db, e := NewDBMaker("test", types).Compute(what); e != nil {
		err = e
	} else {
		//text, _ := Marshal(db)
		//log.Println("db", text)

		m := NewStoryModel(db, types)
		if block, _, e := m.BuildRootCmd("test"); e != nil {
			err = e
		} else {
			text, _ := Marshal(*block)
			log.Println("blocks", text)

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

// generates a noun directive
// directives are used to start describing scripts.
func xTestSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

// generates a ScriptRef statement
// statements are used in callbacks.
func xTestScriptRef(t *testing.T) {
	what := g.The("fish") //.Is("hungry")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("our fish", text)
	}
}

func xTestExists(t *testing.T) {
	what := The("cabinet", Exists()) //IsKnownAs("armoire")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("The cabinet exists.", text)
	}
}

func xTestKnownAs(t *testing.T) {
	what := The("cabinet", IsKnownAs("the armoire"))
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("The cabinet is known as the armoire.", text)
	}
}

func xTestUnderstanding(t *testing.T) {
	what := Understand("feed {{something}}").As("feeding it")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal(`Understand feed {{something}} as feeding it.`, text)
	}
}

// FIX: evetually these snippets should become part of their test suite
// and we run the matcher externally, generically.
// because ideally, our tests would be near to where they are declared.
func xTestIs(t *testing.T) {
	what := g.The("fish").Is("hungry")
	assert := assert.New(t)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("is our fish hungry", text)
	}
}

//
func xTestJoinAll(t *testing.T) {
	assert := assert.New(t)
	what := core.All(g.The("fish").Is("hungry"), g.The("fish food").Is("found"))
	assert.Len(what.(core.AllTrue).Test, 2)
	if text, e := PhraseText(what); assert.NoError(e) {
		assert.Equal("( is our fish hungry, and is our fish food found )", text)
	}
}
