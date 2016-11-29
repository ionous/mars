package script

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

// ListOfItems intended for use in a The() phrase.
type ListOfItems struct {
	list  string
	items []string
}

// And allows the continuation of items
func (my ListOfItems) And(name string) ListOfItems {
	my.items = append(my.items, name)
	return my
}

// In gives the current object the room for its whereabouts
func In(room string) ListOfItems {
	return ListOfItems{"whereabouts", []string{room}}
}

// Contains gives the current supporter the passed prop as contents.
func Supports(prop string) ListOfItems {
	return ListOfItems{"contents", []string{prop}}
}

// Contains gives the current container the passed prop as contents.
func Contains(prop string) ListOfItems {
	return ListOfItems{"contents", []string{prop}}
}

// Possesses gives the current subject the passed prop as inventory.
func Possesses(prop string) ListOfItems {
	return ListOfItems{"inventory", []string{prop}}
}

// Wears gives the current subject the passed article of clothing.
func Wears(prop string) ListOfItems {
	return ListOfItems{"clothing", []string{prop}}
}

// GenFragment implements script.backend Fragment
func (my ListOfItems) GenFragment(src *S.Statements, top Topic) (err error) {
	list := my.list
	for _, item := range my.items {
		fields := S.KeyValueFields{top.Subject.String(), list, item}
		if e := src.NewKeyValue(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}
