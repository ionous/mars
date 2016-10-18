package s

import (
	"github.com/ionous/mars/script/frag"
)

// Is asserts one or more states of one or more enumerations.
// The enumerations must (eventually) be declared for the target's class. ( For example, via AreEither, or AreOneOf, )
func Is(choice string, choices ...string) frag.Select {
	return frag.Select{append(choices, choice)}
}
