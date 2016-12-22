package rt

// Numbers provides an array literal for floats.
type Numbers struct {
	Values []float64
}

func (l Numbers) GetNumberStream(Runtime) (NumberStream, error) {
	return &NumberIt{list: l.Values}, nil
}

type NumberIt struct {
	list []float64
	idx  int
}

func (it *NumberIt) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *NumberIt) GetNext() (ret Number, err error) {
	if !it.HasNext() {
		err = StreamExceeded("Numbers")
	} else {
		ret = Number{it.list[it.idx]}
		it.idx++
	}
	return
}

// Texts provides an array literal for strings.
type Texts struct {
	Values []string
}

func (l Texts) GetTextStream(Runtime) (TextStream, error) {
	return &TextIt{list: l.Values}, nil
}

type TextIt struct {
	list []string
	idx  int
}

func (it *TextIt) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *TextIt) GetNext() (ret Text, err error) {
	if !it.HasNext() {
		err = StreamExceeded("Texts")
	} else {
		ret = Text{it.list[it.idx]}
		it.idx++
	}
	return
}

// References provides an array literal for object ids.
type References struct {
	Values []ObjEval
}

func (l References) GetObjStream(run Runtime) (ObjectStream, error) {
	return &RefIt{run: run, list: l.Values}, nil
}

type RefIt struct {
	run  Runtime
	list []ObjEval
	idx  int
}

func (it *RefIt) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *RefIt) GetNext() (ret Object, err error) {
	if !it.HasNext() {
		err = StreamExceeded("References")
	} else {
		ref := it.list[it.idx]
		if obj, e := ref.GetObject(it.run); e != nil {
			err = e
		} else {
			ret = obj
			it.idx++
		}
	}
	return
}
