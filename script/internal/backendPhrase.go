package internal

import S "github.com/ionous/sashimi/source"

// Source contains the output of script.
type Source struct {
	*S.BuildingBlocks
}

// BackendPhrases are used to build script into source.
type BackendPhrase interface {
	Build(Source) error
}

// UnknownLocation is a stand-in for the file and line of the phrase used to build a statement. MARS: remove this and replace with the proper file and line!
const UnknownLocation = S.Code("unknown")
