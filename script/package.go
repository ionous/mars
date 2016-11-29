package script

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/inbuilt"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
)

func Package() *mars.Package {
	if script == nil {
		script = &mars.Package{
			Name:         "Script",
			Dependencies: mars.Dependencies(inbuilt.Inbuilt()),
			Interfaces:   (*ScriptInterfaces)(nil),
			Commands:     (*ScriptCommands)(nil),
		}
	}
	return script
}

var script *mars.Package

// https://github.com/ungerik/pkgreflect: could be used via go generate
type ScriptInterfaces struct {
	backend.Spec
	backend.Fragment
	internal.ActionRequirements
	internal.ReverseRelation
}

type ScriptCommands struct {
	*backend.SpecList
	*internal.CanDoIt
	*internal.Choices
	*internal.ClassEnum
	*internal.ClassProperty
	*internal.HaveOne
	*internal.HaveMany
	*internal.ImplyingNothing
	*internal.ImplyingOne
	*internal.ImplyingMany
	*internal.DefaultAction
	*internal.Exists
	*internal.PropertyValue
	*internal.KnownAs
	*internal.ScriptSubject
	*internal.ScriptSingular
	*internal.BeforeEvent
	*internal.WhenEvent
	*internal.AfterEvent
	// GoesToFragment
	// ListOfItems
	*internal.ParserPhrase
	*internal.NounPhrase
	*internal.Requires
	*internal.RequiresOnly
	*internal.RequiresTwo
	*internal.RequiresNothing
}
