package blocks

type Span struct {
	Tag      string
	Text     string
	Path     string
	Sep      Separator `json:"-"`
	Children *Block    `json:",omitempty"`
}
