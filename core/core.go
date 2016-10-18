package core

type Core struct {
	// all.go
	*AllTrue
	// any.go
	*AnyTrue
	// boolEval.go
	*Compare
	*Is
	*Not
	*Empty
	*Equals
	*Exists
	// call:
	*CallWithNum
	*CallWithText
	*CallWithRef
	// change:
	*SetNum
	*SetTxt
	*SetRef
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
	*Statements
	*Error
	*Fails
	// gocall.go
	*GoCall
	// numEval
	*AddNum
	// pointsTo
	*PointsTo
	// print
	*PrintNum
	*PrintText
	*PrintLine
	// property
	*Property
	*NumberValueList
	*TextValueList
	*RefValueList
}
