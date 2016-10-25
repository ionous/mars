package internal

// Topic targets a noun and or its type.
type Topic struct {
	Target, Subject string
}

// Fragment phrases appear in "The" phrases.
type Fragment interface {
	BuildFragment(Source, Topic) error
}

// Fragments array
type Fragments []Fragment
