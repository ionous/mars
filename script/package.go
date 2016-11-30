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
	*internal.AfterEvent
	*internal.BeforeEvent
	*internal.CanDoIt
	*internal.Choices
	*internal.ClassEnum
	*internal.ClassProperty
	*internal.DefaultAction
	*internal.Exists
	*internal.HaveOne
	*internal.HaveMany
	*internal.ImplyingNothing
	*internal.ImplyingOne
	*internal.ImplyingMany
	*internal.KnownAs
	*internal.NounPhrase
	*internal.NumberValue
	*internal.ParserPhrase
	*internal.RefValue
	*internal.Requires
	*internal.RequiresOnly
	*internal.RequiresTwo
	*internal.RequiresNothing
	*internal.ScriptSubject
	*internal.ScriptSingular
	*internal.TextValue
	*internal.WhenEvent
}
