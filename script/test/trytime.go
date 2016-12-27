package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

// Trytime adapts a mars runtime to the needs of script testing.
type Trytime interface {
	// Parse executes the passed string as user input;
	// returns the output or error.
	Parse(string) ([]string, error)
	// Parse executes the passed string as user input;
	// returns the output or error.
	Run(string, []meta.Generic) ([]string, error)
	// Execute runs the passed statements;
	// returns the output or error.
	Execute(rt.Statements) ([]string, error)
	// Test the passed boolean eval, returning error if not succesfull.
	Test(rt.BoolEval) error
}
