package script

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"strings"
)

var Directions = []string{"north", "south", "east", "west", "up", "down"}

func _makeOpposites() map[string]string {
	op := make(map[string]string)
	for i := 0; i < len(Directions)/2; i++ {
		a, b := Directions[2*i], Directions[2*i+1]
		op[a], op[b] = b, a
	}
	return op
}

var opposites = _makeOpposites()

// Going begins a statement connecting one room to another via a movement direction.
// Direction: a noun of "direction" type: ex. north, east, south.
func Going(direction string) GoingFragment {
	return GoingFragment{GoesFromFragment{fromDir: direction}}
}

// Through causes directional movement to pass through an explicit departure door.
func (goingFrom GoingFragment) Through(door string) GoesFromFragment {
	goingFrom.fromDoor = door
	return goingFrom.GoesFromFragment
}

// Through begins a statement connecting one room to another via a door.
// Door: The exit from a room.
func Through(door string) GoesFromFragment {
	return GoesFromFragment{fromDoor: door}
}

// ConnectsTo establishes a two-way connection between the room From() and the passed destination.
func (goesFrom GoesFromFragment) ConnectsTo(room string) GoesToFragment {
	return GoesToFragment{from: goesFrom, toRoom: room, twoWay: true}
}

// ArrivesAt establishes a one-way connection between the room From() and the passed destination.
func (goesFrom GoesFromFragment) ArrivesAt(room string) GoesToFragment {
	return GoesToFragment{from: goesFrom, toRoom: room}
}

// Door specifies an optional door to arrive at in the destination room.
func (goesTo GoesToFragment) Door(door string) backend.Fragment {
	goesTo.toDoor = door
	return goesTo
}

type GoingFragment struct {
	GoesFromFragment
}

type GoesFromFragment struct {
	fromDir, fromDoor string
}

// GoesToFragment intended for use in a The() phrase.
type GoesToFragment struct {
	from           GoesFromFragment
	toRoom, toDoor string
	twoWay         bool
}

// GenFragment implements script.backend Fragment
func (goesTo GoesToFragment) GenFragment(src *S.Statements, top backend.Topic) error {
	from := newFromSite(top.Subject.String(), goesTo.from.fromDoor, goesTo.from.fromDir)
	to := newToSite(goesTo.toRoom, goesTo.toDoor, goesTo.from.fromDir)

	s := NewScript(from.makeSite(), to.makeSite())

	// A departure door (has a Understand) arrival door
	s.The(from.door.str, HasText("destination", T(to.door.str)))
	if goesTo.twoWay {
		s.The(to.door.str, HasText("destination", T(from.door.str)))
	}
	// A Room+Travel Direction (has a Understand) departure door
	// ( if you do not have an deptature door, one will be created for you. )

	dir := xDir{goesTo.from.fromDir}
	if dir.isSpecified() {
		s.Add(dir.makeDir())
		s.The(from.room.str, HasText(dir.via(), T(from.door.str)))

		if goesTo.twoWay {
			s.The(to.room.str, HasText(dir.revVia(), T(to.door.str)))
			// FIX? REMOVED dynamic opposite lookup
			// needs s thought as to how new directions could be added
			// perhaps some sort of "dependency injection" where we can add evaluations
			// -- dynmic compiler generators -- as hooks after ( dependent on ) sets of other instances, classes, etc. so those hooks can use model reflection to generate new, non-conflicting, model data -- this is already similar to the idea of onion skins of visual content, hardpoint hooks, etc.
			//_, err = b.The(to.room.str, Has(dir.revRev(), from.door.str))
		}
	}

	return s.Generate(src)
}

// helper to create departure door if needed
func newFromSite(room, door, dir string) xSite {
	gen := door == ""
	if gen {
		door = strings.Join([]string{room, "departure", dir}, "-")
	}
	return xSite{xRoom{room}, xDoor{door, gen}}
}

// helper to create arrival door if needed
func newToSite(room, door, dir string) xSite {
	gen := door == ""
	if gen {
		door = strings.Join([]string{room, "arrival", dir}, "-")
	}
	return xSite{xRoom{room}, xDoor{door, gen}}
}

type xSite struct {
	room xRoom
	door xDoor
}

func (x xSite) makeSite() backend.Spec {
	s := NewScript(
		The("room", Called(x.room.str), Exists()),
		The("door", Called(x.door.str), In(x.room.str), Exists()))
	if x.door.gen {
		s.The(x.door.str, Is("scenery"))
	}
	return s
}

type xRoom struct {
	str string
}

type xDoor struct {
	str string
	gen bool
}

type xDir struct {
	str string
}

func (x xDir) isSpecified() bool {
	return len(x.str) > 0
}

func (x xDir) makeDir() backend.Spec {
	return The("direction", Called(x.str), Exists())
}

func (x xDir) via() string {
	return x.str + "-via"
}
func (x xDir) opposite() string {
	return opposites[x.str]
}

func (x xDir) revVia() string {
	return x.opposite() + "-via"
}

// FIX? REMOVED dynamic opposite lookup ( see comment in MakeStatement )
// func (x xDir) revVia() string {
// 	return x.str + "-rev-via"
// }
