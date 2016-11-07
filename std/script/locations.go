package script

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type ListOfItems struct {
	list  string
	items []string
}

func (my ListOfItems) And(name string) ListOfItems {
	my.items = append(my.items, name)
	return my
}

// FIX: move these into a standard rules extension package?
func In(room string) ListOfItems {
	return ListOfItems{"whereabouts", []string{room}}
}

func Supports(prop string) ListOfItems {
	return ListOfItems{"contents", []string{prop}}
}

func Contains(prop string) ListOfItems {
	return ListOfItems{"contents", []string{prop}}
}

func Possesses(prop string) ListOfItems {
	return ListOfItems{"inventory", []string{prop}}
}

func Wears(prop string) ListOfItems {
	return ListOfItems{"clothing", []string{prop}}
}

func (my ListOfItems) GenFragment(src *S.Statements, top Topic) (err error) {
	list := my.list
	for _, item := range my.items {
		fields := S.KeyValueFields{top.Subject, list, item}
		if e := src.NewKeyValue(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}
