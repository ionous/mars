package blocks

import (
	"github.com/ionous/mars/export"
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

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
