package core

import "github.com/ionous/mars"

// Core contains all of mar's built-in commands and primitives.
var Core = mars.Package{
	Name: "Core",
	// MARS, FIX: move "kinds" declaration to a custom backend script?
	Commands: (*CoreDL)(nil),
}

type CoreDL struct {
	// all.go
	*AllTrue
	// any.go
	*AnyTrue
	// boolEval.go
	*IsNumber
	*IsText
	*IsObject
	*IsState
	*IsNot
	*IsEmpty
	*IsValid
	// change:
	*SetNum
	*SetTxt
	*SetObj
	*ChangeState
	*Named
	// Choose
	*Choose
	*ChooseObj
	*ChooseText
	*ChooseNum
	// Context
	*Context
	*GetNum
	*GetText
	//each
	*ForEachNum
	*ForEachText
	*ForEachObject
	*IfEach
	*EachIndex
	// exec
	*ExecuteList
	*Error
	*Fails
	// gocall.go
	*GoCall
	// numEval
	*AddNum
	// print
	*PrintNum
	*PrintText
	*PrintLine
	// property access
	*PropertyText
	*PropertyNum
	*PropertyRef
	*PropertyTextList
	*PropertyNumList
	*PropertyRefList
}
