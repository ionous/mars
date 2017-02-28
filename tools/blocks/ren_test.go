package blocks

import (
	"bytes"
	// "encoding/json"
	// "github.com/ionous/mars"
	// "github.com/ionous/mars/core"
	// "github.com/ionous/mars/export"
	// "github.com/ionous/mars/script"
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	// "github.com/ionous/sashimi/util/errutil"
	"github.com/stretchr/testify/assert"
	// "log"
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
				if e := inspect.Inspect(types).Visit(path, r, noun); assert.NoError(e) {
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
					if e := inspect.Inspect(types).Visit(path, r, noun); assert.NoError(e) {
						doc := cursor.Document()
						assert.Len(doc, 4, "root and 3 children")
					}
				}
			}
		}
	}
}

func RunManualRules(what interface{}, rules Rules) (ret string, err error) {
	// /, pack ...*mars.Package
	// pack = append([]*mars.Package{std.Std()}, pack...)
	// if types, e := inspect.NewTypes(pack...); e != nil {
	if types, e := inspect.NewTypes(std.Std()); e != nil {
		err = e
	} else {
		var buf bytes.Buffer
		words := NewWordWriter(&buf)
		g := &Generator{
			words, rules, NewDocument(),
		}
		if e := StackRender("root", g, types, what); e != nil {
			err = e
		} else {
			ret = buf.String()
		}
	}
	return
}

func RunEnglishRules(what interface{}) (ret string, err error) {
	if types, e := inspect.NewTypes(std.Std()); e != nil {
		err = e
	} else {
		var buf bytes.Buffer
		en := NewEnglishGenerator(
			types,
			NewWordWriter(&buf),
			NewDocument(),
		)
		if e := en.Render("root", what); e != nil {
			err = e
		} else {
			ret = buf.String()
		}
	}
	return
}

func TestManualSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	rules := Rules{
		// thinking a scoped "maker" for things?
		WriteType("string", FormatString),
		Prepend("NounDirective", "The"),
		Token("NounDirective.Target", "[subject]"),
		Token("NounDirective.Fragments", "[phrases]"),
		Append("NounDirective", "."),
	}
	if text, e := RunManualRules(what, rules); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}

// next up: use the autogenerating rules
func TestAutoSubject(t *testing.T) {
	what := The("cabinet")
	assert := assert.New(t)
	if text, e := RunEnglishRules(what); assert.NoError(e) {
		assert.Equal("The cabinet [phrases].", text)
	}
}
