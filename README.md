
# Mars
a sashimi dl/dsl research project.

ideas based on [scripting with data](http://dev.ionous.net/2013/03/scripting-with-data.html), literate programming ( [inform7](http://inform7.com), [WEB](https://en.wikipedia.org/wiki/WEB) ), text templating (esp. [handlebars](http://handlebarsjs.com)), [inkle](https://github.com/inkle/ink), ...

trying the experiment of making all notes public, updated over time.

Table of Contents
=================

  * [Mars](#mars)
  * [Table of Contents](#table-of-contents)
  * [Tasks](#tasks)
    * [Future](#future)
  * [Parsing](#parsing)
  * [Scripting](#scripting)
    * [Scripting rules](#scripting-rules)
    * [Collapsing the wall between script and code](#collapsing-the-wall-between-script-and-code)
    * [Relations](#relations)
  * [Model Changes](#model-changes)
    * [the model:](#the-model)
  * [Events](#events)
    * [some thoughts](#some-thoughts)
    * [Calling actions from actions:](#calling-actions-from-actions)
    * [Parameter Naming](#parameter-naming)
    * [Actions structures](#actions-structures)
    * [raising and stopping events](#raising-and-stopping-events)
    * [Issues:](#issues)
  * [Backend phrases](#backend-phrases)
  * [Scanning](#scanning)
  * [Verification](#verification)
    * [text templating](#text-templating)
  * [Value packing](#value-packing)

Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc.go)

# Tasks

* move “text” and “num” to “eval” -- possibly disable relations for now. ( see model changes below )
* goal: get to the point where we can run an action and listen to an event.
   * work more on translating the standard lib and test scripts
* compute a sequence list
* print an inventory list with commas on one line.
* add an object query and object list query
* take a hard look at “look up parents” -- how can it be removed? 

    could it be that with “inferences” ( and the interim ) we could inject a body of code to a particular variable and then use that particular variable in the execution. it would become a generator, a normal -- albeit still “hard-coded” structure command part of the dl -- but it could come from the std lib --be part of it.

* imagine a command like “inklabel”, how would we do it. a “node” object? with choices lists? a command that adds them? side note: i almost think with the right set of blocks, “states” ( just variables ), and statements to set ( aka divert, jump ) -- we could embed the inkle conversation system.
* verification entries: checking the type of property references

## Future

* you are going to have to work on defaults eventually
* an attempt at a text parser
* an attempt at a gui

# Parsing

* im imagining a raw text parser that either generates ( creates ) in memory go instances, or translates to .go files -- in either case which compiles ( or saves, gob maybe ) both go-wise, and game-wise. compiling discards the “script” part of the data in favor of “model”, and only keeps the “execute” part of the data.
* im imagining some parsing of comments above structs for expressions for parsing. CanDo{class string} // {class} can ...
* at a first pass, it could possibly rely on go’s compiler/reflection reported errors as it tokenizes or compiles the script
* chains of struct functions might be difficult, but are somewhat like the descriptions of identifiers in other languages ( yak,bison rules? ) -- and there’s some good advantage in them allowing sentences with “branches”
* it could be cool, in the future, to allow a language “description” file -- the same sort of extracted data that would be needed for a gui -- to, in turn, generate go definitions. ( though you'd still have to write the implementations -- so not sure what exact use it all has )

# Scripting
rewrite the script front end with a scripted execution backend -- we already know that the backend will allow commands to pile on changes to the model, so take that and use it. make the scripting section less special. currently, the compiler cant actually run in layers -- but, we could still add to the structure, a bit at a time i think.

focus on MakeStatement -- could potentially keep that, returning from a command. its close to that already.

**what makes it special to be a backend?** always building structures? not using chains? 

**no**, it's that in sashimi v1, “s” is an instance. the shortcuts build into Statements structure immediately. it should be a set of structures that once traversed build. ie. you can save the commands, you can load the commands, you can execute them to produce Statements, you can execute them in different ways ( via different interface implementations ) as needed.

## Scripting rules 
any rules about creating script objects?

keep it similar to what exists? try to forgo shortcuts -- simple shortcuts might be okay, but chains are only going to make writing a text parser more and more difficult. 

we have so much script already, that keeping close to it is probably a good idea. i really like “thing is type” more than “the type called thing” but... for now... maybe we should keep it. we could provide the other later to clean it up.

“photography data is action data with an actor “the photographer”, an actor “the subject”, and a prop “the camera”. to photograph, go say “maybe some day”.

## Collapsing the wall between script and code
currently the script keeps “text” and “number”, “pointer” ( object reference ) and “relation” -- oh goodness -- what if relation entries were actually queries, that’s magic.

once we have “text eval” -- if we can store that instead -- then we can use “text templates” right at hand. i’m vaguely curious how inform does this -- oh, i can guess ( i think ) -- it probably stores the string and parses the string at run-time ( plus or minus caching. ) some of the basics, now, could be pulled in and verified just like the rest. 

( see model changes, below. )

## Relations

* HaveOne -> ObjectEval backed by a Query [ FirstObjectQuery ]
* HaveMany -> ObjectListEval backed by a Query [ ObjectListQuery ]

# Model Changes
things to look at: the compiler (pending class), the model ( and xmodel ), the runtime.

## the model: 
the instance is currently holding name -> `Value`, where `Value` is `interface{}`. sashimi likely use the property type to interpret that in a switch ( ex. in panic value. )

so the options are:

* variant: text eval, num eval, obj eval, lists, 
    its the list that make me think these should be evals and not executes ( ie. concat result, not ambient accumulaton )
    the variant is giving me pause though, why.
    well in the other bits, we have type safety uberalles --
    we dont have a variant "Eval" that returns just any old thing.

options here include: 

* a structure with pointers to the various evals ( mos ) 
* and maps of each eval type ( som )

concerns:    

* i worry about "forgetting" to add something
* i worry about the difficulty of expanding things
* i worry it will look ugly.


```json
"tunnels-1": {
   "id": "tunnels-1",
   "type": "rooms",
   "name": "tunnels-1",
   "values": {
    "kinds-name": "Maintenance Tunnels",
    "kinds-printed-name": "Maintenance Tunnels",
    "rooms-description": "Cobwebs and roots dangle from the ceiling.\nFrom here it looks like you can go west, east, or south ( back the way you came ).",
    "rooms-east-via": "branch-one-1",
    "rooms-maze-num": 0,
    "rooms-maze-suffix": "1",
    "rooms-north-via": "tunnels-1-arrival-south",
    "rooms-scent": "It's obvious whichever bot was supposed to clean these tunnels hasn't been through here recently.",
    "rooms-sound": "Here, more than anywhere, the ship sounds like a living breathing beast.",
    "rooms-west-via": "branch-two-1"
   }
  }
```



# Events
so something strange is starting to happen -- things which might be scripted actions, are becoming dl machines. and when i look at events -- which are they? they are right on the boundary -- 
```go
s.The("actors", Can("give it to").And("giving it to").RequiresOne("actor").AndOne("prop"), To("give it to", execute), Before("giving it to").Always(execute))
```

## some thoughts

> It's almost like auto-generating a struct that is then used by the code:

```go
type GiveItTo struct {
  ID string 
  Name string 
  Event string 
  Nouns []string 
  Actions []string
} 
```

> Unlike dl calls which are statically typed, action calls are typed more softly.

The code, with all mismatched parameters:
```go
g.The(“table”, Go(“give it to”, g.The(“cabinet”, g.The(“player”))
```

currently compiles because there is no chance to check those parameters. Type-checking and error handling happens at run-time instead.

With the inference compiler the calls to the action in the story might someday decide the input types. in the meantime, with the dl, we might want to push type checking to the backend as much as we can. even then, we might still need some run-time type checking or fallbacks.

> We already know we want to support actions containing text

For instance, it sucks that “ScriptSays” and “ActorSays” are special parts of the game interface, which then generate fake actions for the command stream code -- instead of just events caught by the host ( ex. the console, or command stream code. ) It seems like, the connection points “requires one actor” might be able to be evals.

>  The actor action entries seem like they could be a new Function type right alongside num and text and ref in the model. 

There’s some ambiguity here: the dl doesn't have “functions” per se. It has either “commands” ( which are strongly typed per each command, and therefore are a bad match with a single type “Function” ), or it has Execute blocks -- but, execute blocks don’t take parameters. The closest they take is “Context” for looking things ( `GetText{}`, or `GetObject{}` ) by name. 

Before storing an execute block, we could wrap with a context generator which contains the names, the CanDo phrase could add that.

> How do we call such a function? 

Currently it happens from two places:

* player input: `script.Execute(action string, Matching(...) )`
* action callbacks: `IObject.Go(action string, params ...IObject)`

ex. `actor.Go("acquire it", prop)`

the action callbacks are the first one to worry about --
“Instead of showing something to someone, try giving the noun to the second noun.”

when you call a function, you need parameters. how the f can you do that -- you don't know what the types are. `Go{}.Param{}.And{}.And{}`

each on returning a bigger and bigger structure, or since its more inline with the rest, via shortcuts, where `g.The()` returns a shortcut helper. [ this part isn't expected to be translated into the gui well, but would be good for compatibility  i suppose technically, you could have shortcuts that build into a buffer so you could keep the current look of things -- tho if statements would be dangerous ]

`g.The(“object”).Go(“run”, ...)` and then you could evaluate the type in the helper like you did with the `Say()` helper.

Would produce a GoCall{} structure:  
```go
type GoCall struct {
    Action Property ( which includes the ObjEval and a property id ) 
    Parameters Execute[] // look for good ways to limit push evals?
}
```

i suppose we could declare a new interface -- not Execute -- but Push -- i don't really see anything wrong with that. it more or less would stop evilness, and you implement what -- one per eval? thats basically what you do with text -- it just doesn't have a special “Say” interface -- but that could be cool maybe.

The alternative to 3 pointers and type, is some sort of Push per type -- ( accumulation vs. concatenate again )  [ or would it have to be more like an Add to distinguish layers? handlebars has a kind of layers concept ( in that you have to use paths to access parents ) and i think that might be best for now ( rather than leaking parameters through from earlier functions, which would make reusing in different contexts -- and verifying types -- much more difficult ) ] a name-lookup [ much like functions push variables onto the stack ]

The runtime would need the struct and the three pointers still -- or, with golang -- an interface we can just switch on -- which seems okay.

Look at noun, second noun, in inform -- i want to see more the shortcut names

## Calling actions from actions:

* replace `RunWithText` and `SpeechPhrase` with `GoCall`
* figure out a way to mock these in mars as a test -- it wants model, metal. so sorry its huge. can you give it one?

## Parameter Naming
```go 
    s.The("actors",
        Can("give it to").And("giving it to").RequiresOne("actor").AndOne("prop"),
        To("give it to").With("the giver").And("the givee").And("the gift").Go(),
```

* we need defaults 
* replace the script's can/do implementation with an “Aliases” block -- DeclareParameters? -- which encodes the requested names and types;
when a statement asks for a name via `GetObj(“giver”}`, etc. you look at the Aliases block and either match in index order -- re: push -- or add a map, possibly a map so you can cache and not evaluate. in fact -- because of side-effects and timing when you trigger an action, evaluate the passed evals, and assign the results - with type checking -- to the map by name. ie. give the Action an interface or parameters to do this.  -- still a little squirrelly, but it'll be okay i think. you might even start here when prototyping this.
* add a new PropertyType for ExecuteBlock
* remove “actions” as a special map from the script Model
* remove the action.Target, action.Context, etc. give everything names instead; eg. `s.The("actors", Can("acquire it").And("acquiring it").RequiresOne("prop").Meaning(“acquirers”, “acquire”, “acquisitions”)`
* make the names for the reflected actions match; don’t worry about rotation.

## Actions structures
idea 1:
```go
    s.The("kinds", Called("actions"), Exist())
    s.The("actions", Called("photography"),
            Have("photographer", "actor"), 
            Have("subject", "actor"), 
            Have("camera", "prop"))
```

you could maybe assume that order of declaration is order of nouns. one annoying thing, is that reflecting is going to force you to do that again and again. alternatively, you could use the same structure in each. the action data vs. action itself. that might be nice just to avoid the shuffle.  

NOTE: the “the” function is an accumulator, rather than a return value concatenate.

```go
        s.The("kinds", Called("actions data"), WithSingularName(“action data“), Exist())
        s.The("actions data", Called("photography data"),
                Have("photographer", "actor"), 
                Have("subject", "actor"), 
                Have("camera", "prop"))
```

something like `CanDo` would be entirely different -- and it would imply you can have actions, events?, on nothing. There’s a weird bit around the the idea that the execution is stored in the instance -- and that has to match the action first parameter.  `HasTo` is rather rough -- by convention, or by definition maybe; even though other things could be stored there.  

the slot would have to have the parameters set as part of the execution definition -- `ActionProperty { Execute, ObjectEval }` -- and if the object eval doesn't match the execute during call -- its an error. none of that is too terrible.

**we don't have defaults right now, and you would need them**

>actors have action photograph with “photography data”. “photography data” is action data with a photographer actor, a subject actor, and camera prop. actors have default photograph with 


> actors can photograph and photographing requires photography data. photography data is action data with an actor “the photographer”, an actor “the subject”, and a prop “the camera”. to photograph, go say “maybe some day”.

## raising and stopping events

stop after, i don't actually see how you could do this with out the runtime interface -- actually, yes i do: but it'd be a bit squirrelly: you could set a variable, possibly even set a scoped variable so it magically resets each time, stop after would set it -- the gnarly part is the game system would have to check said variable, so its a part of the system that always has to be there. to be fair, you might be pushing t,s,c “object scope” as well ( with custom shortcut names ) -- so that whole bit could be in a sub-interface.
Encoding and Serialization

Every time there is a new “root interface” -- you either need to define a struct containing all of the possible implementations, or you need to serialize via gob and register all of the possible types.

Go already internally has the type on every interface. But, it doesn’t have a list of them. So, to *deserialize* -- even via gob -- registration is required.

You could “register” types by putting them in a structure --- given go rules, i dont even think they need field names just: struct{ *A; *B; *C }  -- and then reflect over that structure for gob. 

The nice thing about this method ( but see  below ) is you can could then serialize to json ( and protobuf ) more readily, and it would give you a place to put “struct tags” if you wanted to write a text-to-go parser. [according to gob zero default rules, empty pointers won’t be encoded]

## Issues:

actually, i don't think we could serialize to json easily --
imagine “Statement” as `struct { *A, *B, *C }` we’d want to be able to store extension functions there too ( ex. from std, facts, quips ) -- but we’d have no way to extend it. auto-generation of a huge bloby structure -- via the ast or imports library -- os one method. another possibility might be something like: MyGameStatements { std.Statements, core.Statements, facts.Statements } and then you'd wrap all of your calls in `G()` which would -- maybe reflect to set -- sounds awful. custom serialization might be better.

for the moment, we wont actually use this structure because statement declaration would become terribly ugly. if we switched to shortcuts, we could possibly put the ugliness in the short cut : rather than returning an interface you would return the structure with the selected item. execution would be a pain though, without marking the type with a switch wed have to walk the whole darn thing field by field. [ you could maybe add a non serializing field to cache that ]

**NOTE**: without actually using this structure -- custom encoding is required to encode the structure name. gob is one custom encoding ( which again needs registration for deserialization. ) the golang default xml and json encoders DO NOT record the concrete type underlying an interface ( xml does in many cases, but not in lists. ) this makes it very difficult ( essentially impossible without error ) to deserialize. 

# Backend phrases
commands so far have had their “front-end” -- how it looks to type them into golang -- and their runtime implementation. but, it’s also possible to give them a backend interface.

this first came up for saying sequences of text: say “a,b,c,d” -- to cycle through those options in order each time you say the text, you need either need some sort of counter, or other sort information on what was last said. basically, you need storage. 

*That’s fascinating because a single statement implying the declaration of both code and data -- and that’s an exceedingly rare pattern in most languages. even in most text templating language -- angular, handlebars, etc. -- when you “range” over a list, you declare the iterator: ng-repeat="item in items".* 

*That said, text templating loops generally do provide implicit loop variables such as $index, $first, $last, $key; but, those are the closest examples I could find. In C++, you can use template meta-programming. In C, you can use macros which use file and line substitutions to generate unique variable names*

Building in that capability in a formal way, however, opens some possibilities. You could, for instance, assert the type of a property access -- and let the compiler detect any conflicts.

# Scanning
note: for “scanning” through the statements for backend implementations,  we actually could just implement the build phase for things like “Statements” which simply walks the statements looking for build phases. In that way, be it for loops, and any other block statements, we can expand them. the alternative would be reflection generically evaluating each structure and array recursively until we find things -- it might be more powerful, if time-consuming, to implement build manually.

# Verification
could do this in a simple way right now, even without the inference compiler, add a new type to BuildingBlocks called Verification, instead of Assertion. Assertions can create -- Verifications don't -- something like that.
sequences

## text templating
see also the text templating doc.

1. maybe it’d be nice to select any sort of list -- the select would have to be separate from the list, and it'd have to work on index or something to be generic with the various list types. 
1. without the “inference compiler” you have to fudge the variable generation -- you'd split the meaning of a token into its runtime behavior and its compile time behavior, possibly by optionally implementing a “compile time” interface ( would that still be needed with the inference, maybe -- maybe you would programmatically add the inference, it does have to come from somewhere ) -  you sketch over the tree searching for the interface(s) -- you'd invoke them to generate some data for the game --
1. you could have a “scope” for variable generation -- but it you'd have to be very careful with it, make it separate from local vars probably, and/or tied particularly to sequence generation: reason being -- i think for a loop where you say a sequence 5 times, youd want to get 5 variations of the sequence and not 5 sequences

# Value packing

needed for literals in lists; basically the issue is:

* in the course of printing we want to say something like GetText{}
* we might have source of that text might be an "object" or "value" but we want to use the same command ( we cant statically guarantee the right command gets used in the right context, so might as well only have one. )
* so, we normalize the value into a "Scope" --
* since Object properties return `Value` ( or `Values`, i guess ) we returned Value from scope - so we needed to push `Value` as well, especially from list who's interface is not `Value` but primitive

options seem to be:

* in the dsl, store `Value`. we cant `Value` is an interface, and i don't think its good to store metal specific things in mars files.
* create a new variant base type used by both mars and metal. i think it needs to e variant for the sake of panic value in meta. the variant may or may not have to include list.
*  pass prims around, `GetText` is currently leveraging the "panic" behavior -- when it really should error -- so switching on prim type manually-- similarly when packing up a value from ObjectScope, translate it.
