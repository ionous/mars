package blocks

type Span struct {
	Tag      string
	Text     string
	Path     string
	Sep      Separator
	Children *Block `json:",omitempty"`
}
