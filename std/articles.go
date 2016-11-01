package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/lang"
)

// TheUpper is equivalent to Inform7's [The]
type TheUpper struct {
	Noun rt.ObjEval
}

// TheLower is equivalent to Inform7's [the]
type TheLower struct {
	Noun rt.ObjEval
}

// AnUpper is equivalent to Inform7's [A/An]
type AnUpper struct {
	Noun rt.ObjEval
}

// TheUpper is equivalent to Inform7's [a/an]
type ALower struct {
	Noun rt.ObjEval
}

func (t TheUpper) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(run, t.Noun, "the"); e != nil {
		err = e
	} else {
		ret = rt.Text(s)
	}
	return
}

func (t TheLower) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(run, t.Noun, "The"); e != nil {
		err = e
	} else {
		ret = rt.Text(s)
	}
	return
}

func (t AnUpper) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(run, t.Noun, ""); e != nil {
		err = e
	} else {
		ret = rt.Text(lang.Capitalize(s))
	}
	return
}

func (t ALower) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if s, e := articleNamed(run, t.Noun, ""); e != nil {
		err = e
	} else {
		ret = rt.Text(s)
	}
	return
}

// You can only just make out the lamp-post.", or "You can only just make out _ Trevor.", or "You can only just make out the soldiers."
func articleNamed(run rt.Runtime, noun rt.ObjEval, article string) (ret string, err error) {
	if obj, e := noun.GetObject(run); e != nil {
		err = e
	} else {
		run := scope.Make(run, scope.NewObjectScope(obj))
		if name, e := (GetText{"printed name"}.GetText(run)); e != nil {
			err = e
		} else if proper, e := (IsState{obj, "proper named"}.GetBool(run)); e != nil {
			err = e
		} else {
			name := name.String()
			if proper {
				ret = lang.Titleize(name)
			} else {
				if len(article) == 0 {
					if indefinite, e := (GetText{"printed name"}.GetText(run)); e != nil {
						err = e
					} else {
						article = indefinite.String()
						if len(article) == 0 {
							if plural, e := (IsState{obj, "plural named"}.GetBool(run)); e != nil {
								err = e
							} else {
								if plural {
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
				// if not, its probably an error.
				if len(article) > 0 {
					ret = article + " " + name
				}
			}
		}
	}
	return
}
