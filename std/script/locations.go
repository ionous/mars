package script

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

// And allows the continuation of items
// func (l *ItemList) And(name string) ListOfItems {
// 	l.Items = append(l.Items, name)
// 	return l
// }

// In gives the current object the room for its whereabouts
func In(noun string) InLocation {
	return InLocation{noun}
}

type InLocation struct {
	Location string `mars:"is in [location];noun"`
}

func (l InLocation) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "whereabouts", []string{l.Location})
}

// Supports gives the current supporter the passed prop as contents.
// In gives the current object the room for its whereabouts
func Supports(nouns ...string) SupportsContents {
	return SupportsContents{nouns}
}

type SupportsContents struct {
	Contents []string `mars:"supports;nouns"`
}

func (l SupportsContents) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "contents", l.Contents)
}

// Contains gives the current container the passed prop as contents.
func Contains(nouns ...string) ContainsContents {
	return ContainsContents{nouns}
}

type ContainsContents struct {
	Contents []string `mars:"contains;nouns"`
}

func (l ContainsContents) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "contents", l.Contents)
}

// Possesses gives the current subject the passed prop as inventory.
func Possesses(nouns ...string) PossessesInventory {
	return PossessesInventory{nouns}
}

type PossessesInventory struct {
	Inventory []string `mars:"has;nouns"`
}

func (l PossessesInventory) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "inventory", l.Inventory)
}

// Wears gives the current subject the passed article of clothing.
func Wears(nouns ...string) WearsClothing {
	return WearsClothing{nouns}
}

type WearsClothing struct {
	Clothing []string `mars:"wears;nouns"`
}

func (l WearsClothing) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "clothing", l.Clothing)
}

// genList implements script.backend Fragment
func genList(src *S.Statements, top Topic, list string, items []string) (err error) {
	for _, item := range items {
		fields := S.KeyValueFields{top.Subject, list, item}
		if e := src.NewKeyValue(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return
}
