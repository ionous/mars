package rt

import "github.com/ionous/sashimi/meta"

type Scope interface {
	FindValue(string) (meta.Generic, error)
	ScopePath() []string
}
