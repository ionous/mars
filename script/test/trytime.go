package test

import "github.com/ionous/mars/rt"

// Trytime adapts a mars runtime to the needs of script testing.
type Trytime interface {
	// Parse executes the passed string as user input;
	// returns the output or error.
	Parse(string) (string, error)
	// Execute runs the passed statement;
	// returns the output or error.
	Execute(rt.Execute) (string, error)
	// Test the passed boolean eval, returning error if not succesfull.
	Test(rt.BoolEval) error
}
