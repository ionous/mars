package script

import "github.com/ionous/mars/script/backend"

func Script(specs ...backend.Spec) backend.Script {
	return backend.Script(specs)
}
