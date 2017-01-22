package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
	. "github.com/ionous/mars/std/script"
	"github.com/ionous/sashimi/util/errutil"
)

func init() {
	m := NewScript()
	// 1. A Room (contains) Doors
	m.The("openers",
		Called("doors"),
		Exist())

	// 2. A Departure Door (has a Understand) Arrival Door
	m.The("doors",
		// exiting using a door leads to one location
		HaveOne("destination", "door").
			// one door can be the target of many other doors
			Implying("doors", HaveMany("sources", "doors")),
	)

	// 3. A Room+Travel Direction (has a Understand) Departure door
	// FIX: without relation by value we have to give each room a set of explict directions
	// each direction relation points to the Understand door
	for _, dir := range Directions {
		// moving in a direction, takes us to a room's entrance.
		m.The("rooms", HaveOne(dir+"-via", "door").
			// FIX: the reverse shouldnt be required; something in the compiler.
			Implying("doors", HaveMany("x-via-"+dir, "rooms")))
		// FIX? REMOVED dynamic opposite lookup
		// // the reverse directions are necessary:
		// // we dont know the set of all directions at compile time
		// // ( we have the default set, but users could add more ).
		// The("rooms", HaveOne(dir+"-rev-via", "door").
		// 	Implying("doors", HaveMany("x-rev-via-"+dir, "rooms")))

		// east is known as "e"
		m.The(dir, IsKnownAs(dir[:1]))
	}

	// Directions:
	m.The("kinds", Called("directions"),
		HaveOne("opposite", "direction").
			//FIX: the reverse shouldnt be required; something in the compiler.
			Implying("directions", HaveOne("x-opposite", "direction")),
	)
	for i := 0; i < len(Directions)/2; i++ {
		a, b := Directions[2*i], Directions[2*i+1]
		m.The("direction", Called(a), HasRef("opposite", b))
		m.The("direction", Called(b), HasRef("opposite", a))
	}

	m.The("actors",
		Can("go to").And("going to").RequiresOnly("direction"),
		To("go to",
			// try the forward direction:
			Using{
				// north-via is a relation,
				Object: DoorHack{
					g.The("actor").Object("whereabouts"),
					g.The("action.Target")},
				Run:  g.Go(g.The("actor").Go("go through it", g.TheObject())),
				Else: g.Go(g.Say("You can't move that direction.")),
			},
		))
	m.The("actors",
		Can("go through it").And("going through it").RequiresOnly("door"),
		To("go through it", g.ReflectToTarget("be passed through")),
	)
	m.The("doors",
		Can("be passed through").And("being passed through").RequiresOnly("actor"),
		To("be passed through",
			Using{
				// the destination of a door is another door
				Object: g.The("door").Object("destination"),
				Run: g.Go(
					Using{
						// the whereabouts of the door, is the room
						Object: g.TheObject().Object("whereabouts"),
						Run: g.Go(
							Choose{
								If: g.The("door").Is("closed"),
								True: g.Go(
									Choose{
										If:    g.The("door").Is("locked"),
										True:  g.Go(g.The("door").Go("report locked", g.The("actor"))),
										False: g.Go(g.The("door").Go("report currently closed", g.The("actor"))),
									}),
								False: g.Go( // FIX: player property change?
									// at the very least a move action.
									Move(g.The("actor")).To(g.TheObject()),
									g.TheObject().Go("report the view")),
							}),
					}),
			},
		),
		Can("report currently closed").
			And("reporting currently closed").
			RequiresOnly("actor"),
		To("report currently closed",
			// FIX: g.The("actor").Says("It's closed."),
			g.Say("It's closed."),
		))
	// understandings:
	// FIX: noun Understand: so that >go north; >go door. both work.
	// FIX: noun alias: Understand "n" as north.
	m.Understand("go {{something}}").As("go to")
	m.Understand("enter {{something}}").As("go through it")
	pkg.Add("Movement", m)

	s := NewScript()
	//exit door and its room, with optional door
	s.The("player", Exists(), In("the lobby"))

	s.The("room", Called("the lobby"),
		// two-way direction
		Going("up").Through("the trap door").ConnectsTo("the parapet"),
		// one-way directions
		Going("down").ArrivesAt("the basement"),
	)
	s.The("foyer",
		// direction to room, reverses
		Going("north").ConnectsTo("the outside"),
		// direction to room, no-reverse
		Going("west").ArrivesAt("the lobby"),
	)
	s.The("lobby",
		// non-commensurate directions
		Going("north").ArrivesAt("the foyer"))
	// explicitly declaring the door should be fine.
	s.The("door", Called("the cellar door"), Exists())
	// direction to door, does not reverse
	s.The("basement", Going("south").
		ArrivesAt("the outside").Door("the cellar door"),
	)
	// not explicitly declaring the door should also work:
	//     The("door", Called("the cellar door"), Exists())
	// door-to-door two-way.
	s.The("foyer", Through("the curtain").
		ConnectsTo("the cloakroom").Door("the cloakroom-curtain"),
	)
	// FIX: want to map "name" to a property, and if it doesn't exist fall back on split id.
	// FIX? wonder if you could inject a report of some kind to pull in the description /brief of a door
	// automatically to match it's other side.
	s.The("door", Called("curtain"), HasText("brief", T("A red velvet curtain.")))
	s.The("door", Called("cloakroom-curtain"), HasText("brief", T("A red velvet curtain.")))

	move := func(cmd, dest string) test.Trial {
		return test.Parse(cmd).
			Expect(g.The("player").Object("whereabouts").Equals(g.The(dest))).
			Else(g.Say(g.The("player").Object("whereabouts")))
	}

	// test moving around
	pkg.AddTest("Moving",
		test.SetupScript(s).Try("moving about",
			test.Expect(g.The("player").Object("whereabouts").Equals(g.The("lobby"))).
				Else(g.Say(g.The("player").Object("whereabouts"))),
			move("go west", "Lobby").Match("You can't move that direction."),
			move("go east", "Lobby").Match("You can't move that direction."),
			move("go up", "Parapet"),
			move("go down", "Lobby"), // first two way
			move("go down", "Basement"),
			move("go up", "Basement").Match("You can't move that direction."),
			move("go south", "Outside"),
			move("go south", "Foyer"),
			move("enter curtain", "Cloakroom"),
		),
	)
}

type DoorHack struct {
	Room, Direction rt.ObjEval
}

// given a room and a direction of movement, we need to find the door to use
func (dh DoorHack) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if room, e := dh.Room.GetObject(run); e != nil {
		err = errutil.New("door hack, failed to get room", e)
	} else {
		// maybe some doors dont have destinations
		if room.Exists() {
			if dir, e := dh.Direction.GetObject(run); e != nil {
				err = errutil.New("door hack, failed to get direction", e)
			} else {
				// north-via
				relName := string(dir.GetId()) + "-via"
				if doorRelation, ok := room.FindProperty(relName); !ok {
					err = errutil.New("door hack, failed to find relation", room, relName)
				} else {
					v := doorRelation.GetGeneric()
					if eval, ok := v.(rt.ObjEval); !ok {
						err = errutil.New("door hack, failed to get relation eval", v)
					} else {
						ret, err = eval.GetObject(run)
					}
				}
			}
		}
	}
	return
}
