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
	*ClassEnum
	*DefaultAction
	*Exists
	*PropertyValue
	*KnownAs
	*ScriptSubject
}
