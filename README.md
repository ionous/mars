
# Mars
a sashimi dl research project.

a [data language](http://dev.ionous.net/2013/03/scripting-with-data.html)(DL) combines some aspects of functional programming with the "gang of four" [command pattern](https://en.wikipedia.org/wiki/Command_pattern). 

a dl leverages the idea that a function can be represented as a class ( with the execution semantics and return value encapsulated in an interface the class implements ); and, uses the fact that simple ( [POD-like](https://en.wikipedia.org/wiki/Passive_data_structure) ) classes can be easily de/serialized. in mars, these two facts combine to allow a sort of "abstract syntax tree" which can be loaded from disk using off-the-shelf libraries, and which can then be executed directly without the need for additional parsing steps.

in its original conception there would be a gui-based tool which would allow the creation of graphical code. in this current implementation, the mars dl is embedded inside of golang using struct literals as commands, and "shortcut" functions which act as macros ( returning aggregates of structs. ) 

the mars dl acts less like a [domain-specific-language](https://en.wikipedia.org/wiki/Domain-specific_language) and more like a concrete syntax tree for a compiler's front-end.  

for dsls which are embedded in a host language, the dsl code normally gets executed near in time to when the dsl functions are invoked. while for dsls which are defined externally to a host language, the host language typically invokes an interpreter or parser to execute them. 

with mars, the host's compiler (go) generates the command tree from the statements embedded in the language (go source code). and, then later, a separate program can load the command tree to execute those statements. 

although mars is currently embedded in go, other languages ( including a graphical ui ) can act as compilers for the dl, so long as the complete set of command structures and fields are (re)declared. other languages can also host execution, so long as the command interfaces are implemented.

this is a work in progress. if you're interested, the [wiki](https://github.com/ionous/mars/wiki/Development-Notes) has more info on its progress.