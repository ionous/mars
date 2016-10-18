package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/lang"
)

// TheUpper is equivalent to Inform7's [The]
type TheUpper struct {
	Noun rt.RefEval
}

// TheLower is equivalent to Inform7's [the]
type TheLower struct {
	Noun rt.RefEval
}

// AnUpper is equivalent to Inform7's [A/An]
type AnUpper struct {
	Noun rt.RefEval
}

// TheUpper is equivalent to Inform7's [a/an]
type ALower struct {
	Noun rt.RefEval
}

func (t TheUpper) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(r, t.Noun, "the"); e != nil {
		err = e
	} else {
		ret = rt.Text(s)
	}
	return
}

func (t TheLower) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(r, t.Noun, "The"); e != nil {
		err = e
	} else {
		ret = rt.Text(s)
	}
	return
}

func (t AnUpper) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(r, t.Noun, ""); e != nil {
		err = e
	} else {
		ret = rt.Text(lang.Capitalize(s))
	}
	return
}

func (t ALower) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(r, t.Noun, ""); e != nil {
		err = e
	} else {
		ret = rt.Text(s)
	}
	return
}

// You can only just make out the lamp-post.", or "You can only just make out _ Trevor.", or "You can only just make out the soldiers."
func articleNamed(r rt.Runtime, noun rt.RefEval, article string) (ret string, err error) {
	if ref, e := noun.GetReference(r); e != nil {
		err = e
	} else if n, e := MakeObject(r, ref); e != nil {
		err = e
	} else if printed, ok := n.FindProperty("printed name"); !ok {
		err = errutil.New("object doesnt have printed names?")
	} else {
		choice := MakeStringId("proper-named")
		if proper, ok := n.GetPropertyByChoice(choice); !ok {
			err = errutil.New("object doesnt have proper names?")
		} else {
			name := printed.GetValue().GetText()
			if choice == proper.GetValue().GetState() {
				ret = lang.Titleize(name)
			} else {
				if len(article) == 0 {
					if p, ok := n.FindProperty("indefinite article"); !ok {
						err = errutil.New("object doesnt have indefinite articles?")
					} else {
						article = p.GetValue().GetText()
						if len(article) == 0 {
							choice := MakeStringId("plural-named")
							if plural, ok := n.GetPropertyByChoice(choice); !ok {
								err = errutil.New("object doesnt have plural named?")
							} else {
								if choice == plural.GetValue().GetState() {
									article = "some"
								} else if lang.StartsWithVowel(name) {
									article = "an"
								} else {
									article = "a"
								}
							}
						}
					}
				}
				if len(article) > 0 {
					ret = article + " " + name
				}
			}
		}
	}
	return
}
