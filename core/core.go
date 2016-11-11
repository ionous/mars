package core

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
)

// Core contains all of mar's built-in commands and primitives.
func Core() *mars.Package {
	if core == nil {
		core = &mars.Package{
			Name: "Core",
			// MARS, FIX: move "kinds" declaration to a custom backend script?
			Scripts:  scripts,
			Tests:    tests,
			Imports:  nil,
			Commands: (*CoreDL)(nil),
		}
	}
	return core
}

var core *mars.Package

var scripts mars.SpecList

func addScript(_ string, specs ...backend.Spec) {
	scripts = append(scripts, backend.SpecList(specs))
}

var tests []test.Suite

func addTest(name string, units ...test.Unit) {
	tests = append(tests, test.NewSuite(name, units...))
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
	*Using
	*GetNum
	*GetText
	*GetObject
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
