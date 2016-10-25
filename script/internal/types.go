package internal

type Types struct {
	// phrases:
	*EventPhrase
	*NounPhrase
	*ParserPhrase
	// fragments ( part of a noun phrase )
	*ActionAssertion
	*Choices
	*ClassProperty
	*DefaultAction
	*Exists
	*KeyValue
	*ScriptSubject
}
