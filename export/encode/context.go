package encode

type Context struct {
	PackageMap
	*TypeRecoder
}

func NewContext() *Context {
	return &Context{
		make(PackageMap),
		NewTypeRecoder(),
	}
}
