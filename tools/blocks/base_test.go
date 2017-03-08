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
func TestFormat(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(DataToString(true), "true")
	list := []interface{}{true, "blue", 5.0}
	assert.Equal(DataToString(list), "true, blue, 5")
	assert.Equal(DataItemAndItem(list), "true and blue and 5")
	assert.Equal(DataItemOrItem(list), "true or blue or 5")
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
				u := p.Usage()
				assert.True(u.IsPrim(), u)
			}
			if p, ok := cmd.FindParam("Fragments"); assert.True(ok) {
				u := p.Usage()
				assert.True(u.IsArray(), u)
			}
		}
		if cmd, ok := types["ScriptRef"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Obj"); assert.True(ok) {
				u := p.Usage()
				assert.True(u.IsCommand(), u)
			}
		}
	}
	if types, e := inspect.NewTypes(export.Export()); assert.NoError(e) {
		if cmd, ok := types["Library"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Types"); assert.True(ok) {
				u := p.Usage()
				assert.True(u.IsBlob(), u)
			}
		}
	}
	if types, e := inspect.NewTypes(script.Package()); assert.NoError(e) {
		if cmd, ok := types["ParserDirective"]; assert.True(ok) {
			if p, ok := cmd.FindParam("Input"); assert.True(ok) {
				u := p.Usage()
				assert.True(u.IsArray(), u)
			}
		}
	}
}
