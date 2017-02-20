package encode

type CommandWalker interface {
	NewCommand(*ParamInfo, *CommandType) (ArgWalker, error)
}

type ArgWalker interface {
	NewCommand(*ParamInfo, *CommandType) (ArgWalker, error)
	NewArray(*ParamInfo, *CommandType) (CommandWalker, error)
	NewValue(*ParamInfo, interface{}) error
}
