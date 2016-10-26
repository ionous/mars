package core

type Core struct {
	// all.go
	*AllTrue
	// any.go
	*AnyTrue
	// boolEval.go
	*IsNumber
	*IsText
	*IsSame
	*IsObject
	*IsNot
	*IsEmpty
	*IsValid
	// change:
	*SetNum
	*SetTxt
	*SetObj
	*ChangeState
	// Choose
	*Choose
	*ChooseRef
	*ChooseText
	*ChooseNum
	// Context
	*Context
	*GetNum
	*GetText
	//each
	*EachNum
	*EachText
	*EachObj
	*IfEach
	*EachIndex
	// exec
	*Executes
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
	*TextProperty
	*NumProperty
	*RefProperty
	*TextListProperty
	*NumListProperty
	*RefListProperty
}
