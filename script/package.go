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
	backend.Directive
	backend.Fragment
	internal.ActionRequirements
	internal.EventTiming
	internal.ParserInput
	internal.ReverseRelation
}

type ScriptCommands struct {
	*internal.AfterEvent
	*internal.BeforeEvent
	*internal.CanDoIt
	*internal.Choice
	*internal.ClassEnum
	*internal.ClassProperty
	*internal.DefaultAction
	*internal.Exists
	*internal.HandleEvent
	*internal.HaveOne
	*internal.HaveMany
	*internal.ImplyingNothing
	*internal.ImplyingOne
	*internal.ImplyingMany
	*internal.KnownAs
	*internal.MatchString
	*internal.NounDirective
	*internal.NumberValue
	*internal.ParserDirective
	*internal.RefValue
	*internal.Requires
	*internal.RequiresOnly
	*internal.RequiresTwo
	*internal.RequiresNothing
	*internal.ScriptSingular
	*internal.ScriptSubject
	*internal.TextValue
	*internal.WhenEvent
}
