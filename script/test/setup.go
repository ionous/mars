package test

import "github.com/ionous/mars/script/backend"

func NewSuite(name string, units ...Unit) Suite {
	return Suite{name, units}
}

func Setup(setup ...backend.Spec) Unit {
	return Unit{setup, nil}
}

func (u Unit) Try(trials ...Trial) Unit {
	u.Trials = append(u.Trials, trials...)
	return u
}
