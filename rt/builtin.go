package rt

type BuiltIn struct {
	*Number
	*Text
	*Reference
	*Numbers
	*Texts
	*References
	// not: objects cannot be stored.
}
