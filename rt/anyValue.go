package rt

import "github.com/ionous/sashimi/meta"

type AnyValue interface {
	GetValue() meta.Generic
}
