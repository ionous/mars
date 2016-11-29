package core

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core/stream"
	"github.com/ionous/mars/inbuilt"
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
)

// Core contains all of mar's built-in commands and primitives.
func Core() *mars.Package {
	if core == nil {
		core = &mars.Package{
			Name: "Core",
			// MARS, FIX: move "kinds" declaration to a custom backend script?
			Scripts:      scripts,
			Tests:        tests,
			Dependencies: mars.Dependencies(inbuilt.Inbuilt(), script.Package()),
			Commands:     (*CoreCommands)(nil),
			Interfaces:   (*CoreInterfaces)(nil),
		}
	}
	return core
}

var core *mars.Package

var scripts backend.SpecList

func addScript(_ string, specs ...backend.Spec) {
	scripts.Specs = append(scripts.Specs, specs...)
}

var tests []test.Suite

func addTest(name string, units ...test.Unit) {
	tests = append(tests, test.NewSuite(name, units...))
}

type CoreInterfaces struct {
	CompareTo
}

type CoreCommands struct {
	// all.go
	*AllTrue
	// any.go
	*AnyTrue
	// boolEval.go
	*IsNum
	*IsText
	*IsObj
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
	*GetObj
	//each
	*ForEachNum
	*ForEachText
	*ForEachObj
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
	*PrintObj
	*PrintLine
	// property access
	*PropertyText
	*PropertyNum
	*PropertyRef
	*PropertyTextList
	*PropertyNumList
	*PropertyRefList
	//
	*stream.ClassStream
	*stream.First
	*stream.MakeStream
	*stream.KeySort
}
