package rt

// Statements is a block of single statement declarations.
// NOTE: it is not a command in and of itself, but only used by other commands. this intentionally excludes anonymous blocks from scripts to decrease noise for scripters.
type Statements []Execute

func MakeStatements(calls ...Execute) Statements {
	return Statements(calls)
}

func (x Statements) Empty() bool {
	return len(x) == 0
}

func (x Statements) ExecuteList(run Runtime) (err error) {
	for _, s := range x {
		if e := s.Execute(run); e != nil {
			err = e
			break
		}
	}
	return err
}
