package inspect

type Arguments interface {
	NewCommand(path Path, baseType, cmdType *CommandInfo) (Arguments, error)
	NewArray(path Path, baseType *CommandInfo, count int) (Elements, error)
	NewValue(path Path, data interface{}) error
}

type Elements interface {
	NewElement(path Path, cmdType *CommandInfo) (Arguments, error)
}
