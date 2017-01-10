package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/lang"
	"strings"
)

type ScriptSubject struct {
	Subject string `mars:"called [subject]"` // name of the class or instance being declared
}

func (c ScriptSubject) GenFragment(src *S.Statements, top Topic) error {
	// FIX: this is only half measure --
	// really it needs to split into words, then compare the first article.
	name := strings.TrimSpace(top.Subject)
	article, bare := lang.SliceArticle(name)
	opt := map[string]string{
		"article":   article,
		"long name": name,
	}
	fields := S.AssertionFields{top.Target, bare, opt}
	return src.NewAssertion(fields, S.UnknownLocation)
}

type ScriptSingular struct {
	Singular types.NamedClass `mars:"has singular name [subject]"`
}

func (c ScriptSingular) GenFragment(src *S.Statements, top Topic) error {
	name := top.Subject
	_, bare := lang.SliceArticle(name)
	opt := map[string]string{
		"singular name": c.Singular.String(),
	}
	fields := S.AssertionFields{top.Target, bare, opt}
	return src.NewAssertion(fields, S.UnknownLocation)
}
