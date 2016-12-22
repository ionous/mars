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
	Location string `mars:";noun"`
}

func (l InLocation) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "whereabouts", []string{l.Location})
}

// Supports gives the current supporter the passed prop as contents.
// In gives the current object the room for its whereabouts
func Supports(noun string) SupportsContents {
	return SupportsContents{[]string{noun}}
}

type SupportsContents struct {
	Contents []string `mars:";nouns"`
}

func (l SupportsContents) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "contents", l.Contents)
}

// Contains gives the current container the passed prop as contents.
func Contains(noun string) ContainsContents {
	return ContainsContents{[]string{noun}}
}

type ContainsContents struct {
	Contents []string `mars:";nouns"`
}

func (l ContainsContents) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "contents", l.Contents)
}

// Possesses gives the current subject the passed prop as inventory.
func Possesses(noun string) PossessesInventory {
	return PossessesInventory{[]string{noun}}
}

type PossessesInventory struct {
	Inventory []string `mars:";nouns"`
}

func (l PossessesInventory) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "inventory", l.Inventory)
}

// Wears gives the current subject the passed article of clothing.
func Wears(noun string) WearsClothing {
	return WearsClothing{[]string{noun}}
}

type WearsClothing struct {
	Clothing []string `mars:";nouns"`
}

func (l WearsClothing) GenFragment(src *S.Statements, top Topic) error {
	return genList(src, top, "clothing", l.Clothing)
}

// genList implements script.backend Fragment
func genList(src *S.Statements, top Topic, list string, items []string) (err error) {
	for _, item := range items {
		fields := S.KeyValueFields{top.Subject.String(), list, item}
		if e := src.NewKeyValue(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return
}
