package uniform

import "github.com/ionous/mars/tools/inspect"

type Context struct {
	inspect.PackageMap
	*inspect.TypeRecoder
}

func NewContext() *Context {
	return &Context{
		make(inspect.PackageMap),
		inspect.NewTypeRecoder(),
	}
}
