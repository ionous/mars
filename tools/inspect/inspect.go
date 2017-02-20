package inspect

type Elements interface {
	NewCommand(*ParamInfo, *CommandInfo) (Arguments, error)
}

type Arguments interface {
	NewCommand(*ParamInfo, *CommandInfo) (Arguments, error)
	NewArray(*ParamInfo, *CommandInfo) (Elements, error)
	NewValue(*ParamInfo, interface{}) error
}
