package frag

import (
	"github.com/ionous/mars/script"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/lang"
	"strings"
)

type SetTopic struct {
	Subject  string // name of the class or instance being declared
	Singular string // optional singular version of that name
}

func (c SetTopic) WithSingularName(name string) Fragment {
	c.Singular = name
	return c
}

func (c SetTopic) Build(src script.Source, top Topic) error {
	// FIX: this is only half measure --
	// really it needs to split into words, then compare the first article.
	name := strings.TrimSpace(top.Subject)
	article, bare := lang.SliceArticle(name)
	opt := map[string]string{
		"article":       article,
		"long name":     name,
		"singular name": c.Singular,
	}
	fields := S.AssertionFields{top.Target, bare, opt}
	return src.NewAssertion(fields, script.Unknown)
}
