package s

import (
	"github.com/ionous/mars/script/frag"
)

// Called asserts the existence of a class or instance.
// For example, The("room", Called("home"))
func Called(subject string) frag.SetTopic {
	return frag.SetTopic{Subject: subject}
}

// Exists is an optional frag for making otherwise empty statements more readable.
// For example, The("room", Called("parlor of despair"), Exists())
func Exists() frag.Fragment {
	return frag.Exists{}
}
func Exist() frag.Fragment {
	return frag.Exists{}
}
