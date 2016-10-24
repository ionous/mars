package rt

import "github.com/ionous/sashimi/meta"

type FindValue interface {
	FindValue(string) (meta.Generic, error)
	ScopePath() []string
}
