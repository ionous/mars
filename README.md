
# Mars
a sashimi data language research project.

# What is a data language?

a [data language](http://dev.ionous.net/2013/03/scripting-with-data.html) (DL) combines some aspects of functional programming with the "gang of four" [command pattern](https://en.wikipedia.org/wiki/Command_pattern). 

a DL leverages the idea that any function can be represented as a class, with the function's body and return value encapsulated in an interface the class implements; and, it uses the fact that simple ( [POD-like](https://en.wikipedia.org/wiki/Passive_data_structure) ) classes can be easily de/serialized. a DL combines these concepts to produce an "abstract syntax tree" which can not only be serialized to and from disk using off-the-shelf tools, but which can be executed directly, without the need for additional parsing steps.

in its original conception there would be a web-based user interface which would allow the authoring of command trees in a graphical manner. in this current implementation, the mars DL is embedded inside of golang using struct literals as commands, and "shortcut" functions which act as macros ( returning aggregates of structs. ) 

# A DL? Don't you mean DSL?

a data language provides a simple ( and possibly unique ) way to implement [domain-specific-languages](https://en.wikipedia.org/wiki/Domain-specific_language) (DSLs).  

broadly speaking, DSL implementations can be separated into [two categories](http://martinfowler.com/books/DSL.html): mini-languages which are *embedded* in a host language. or, mini-languages which are stored *externally*. 

**embedded DSLs** -- of which [LINQ](https://en.wikipedia.org/wiki/Language_Integrated_Query) is one example, but so are custom, one-off [fluent interfaces](http://martinfowler.com/bliki/FluentInterface.html) -- provide commands which are executed near in time to when the DSL functions are invoked ( typically within the same process instance. )

**external DSLs** provide commands which are stored outside of the host language. the host language invokes an interpreter or implements a parser to execute those statements.

note, in both cases, DSLs differ from general purpose languages. domain specific languages can provide very high level commands targeted for a specific application. for instance, in excel, there is a macro language which includes commands to sum whole columns of numbers. that sort of command would make no sense as an instruction built into a language like C, as C lacks a native concept of spreadsheets and columns. ( thank goodness. )

with **data languages**, the host's compiler (in this case [Go](golang.org)) generates commands from statements embedded in the host's language. later, a separate program (in this case, also Go), can load and execute those commands. in this way a DL -- while not necessarily suitable for all purposes -- shares aspects of both types of embedded and external DSLs, giving it some "best of both worlds" benefits.

using compiler terminology, one way of looking at data languages is that the host language provides the DL's concrete syntax tree, the host's compiler provides the DL's compiler front end, and a host's runtime provides the DL's execution environment. ( note: the runtime can also provide a compiler back end: creating storage for variables and generated code which can be baked back into the final data script. )

in short, data languages make building domain specific languages easy.

# Go

Go provides a great host language for DLs. here, [in just a few lines](https://play.golang.org/p/bAkwItKEFH), is a dsl to let you add two integers:

```
func main() {
  script := AddInt{Int{5}, Int{3}}
  fmt.Println("Add two numbers", script.GetInt())
}
```

it has the following type definitions:
```go
type IntEval interface {
  GetInt() int
}
type AddInt struct {
  Augend, Addend  IntEval
}
func (op AddInt) GetInt() int {
  return op.Augend.GetInt() + op.Addend.GetInt()
}
type Int struct {
  Int int
}
func (l Int) GetInt() int {
  return l.Int
}
```
note how even integer literals are represented as commands. this literals to be used any place a function is needed, and allows functions stand in for any literal.

```
func main() {
  script := AddInt{AddInt{Int{1}, Int{4}}, AddInt{Int{1}, Int{2}}}
  fmt.Println("Add several numbers", script.GetInt())
```

to save and load this language, we can use [package gob](https://golang.org/pkg/reflect). gobs have two limitations. first, we must pre-register all the possible interface implementations. second, we must wrap our script in a struct to force the encoder to save the top level of the script as an interface.

here's one easy way to register our types, using [package reflect](https://golang.org/pkg/reflect):
```go
// List all commands here
type IntDSL struct {
  *Int
  *AddInt
}
// Use reflection to register them with gob
t := reflect.TypeOf(IntDSL{})
for i := 0; i < t.NumField(); i++ {
  f := t.Field(i)
  v := reflect.New(f.Type.Elem())
  gob.RegisterName(f.Name, v.Elem().Interface())
}
```

and here's our script wrapper, with an example of save/load:
```go
  type Program struct {
    IntEval
  }
  script := AddInt{Int{5}, Int{3}}
  if e := enc.Encode(Program{script}); e != nil {
    panic(e)
  }
  var p Program
  if e := dec.Decode(&p); e != nil {
    panic(e)
  }
  fmt.Println("Add two numbers again", p.GetInt())
```

the complete example can be found on the [playground](https://play.golang.org/p/1idrsPuIuM). 

note: although this example uses go, almost any language can act as a compiler for a DL. in fact, it's fairly easy to port data languages from one host to another: allowing you to develop tools in one language, while hosting your runtime in another. ( ie. great for games. ) 

# On Mars

mars is a data language for [sashimi](https://github.com/ionous/sashimi). and, sashimi is an interactive fiction engine inspired by literate programming ( [inform7](http://inform7.com) in particular. ) 

sashimi's current scripting system has both declarative and imperative features. its declarative features are implemented as a traditional embedded dsl. its imperative features are written as normal go code. a story file consists of data generated by the dsl, and source code snippets extracted via the ast. 

the goal of moving to mars is to unify sashimi's declarative and imperative features into a single data language, allowing story files to be treated as pure data. 

this means:
* a single executable can support multiple stories;
* iteration time for creating new stories is decreased;
* text templates, and even full conversation trees can be stored and processed in the same manner as any other part of the script;
* tasks like verifying all named objects and properties exist, or shuffling the order of things to say during play can be handled offline in a "pre-compiler" step via the back end;
* with additional work, stories can be authored using text scripts or gui tools.

the [wiki](https://github.com/ionous/mars/wiki/Development-Notes) has more info.