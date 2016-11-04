package test

import "github.com/ionous/mars/script/backend"

func Setup(setup ...backend.Spec) backend.Script {
	return setup
}

func Trials(trials ...Trial) []Trial {
	return trials
}
