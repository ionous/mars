package internal

type Exists struct{}

func (Exists) BuildFragment(Source, Topic) error {
	return nil
}
