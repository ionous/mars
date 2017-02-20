package internal

import (
	"github.com/ionous/sashimi/util/errutil"
	r "reflect"
	"strings"
)

type InterfaceRecord struct {
	name string
	face r.Type
}
type Interfaces []InterfaceRecord

func NewInterface(face r.Type) InterfaceRecord {
	name := face.Name()
	return InterfaceRecord{name, face}
}

func (faces Interfaces) String() string {
	str := make([]string, len(faces))
	for i, s := range faces {
		str[i] = s.name
	}
	return strings.Join(str, ",")
}

func (faces Interfaces) Contains(s r.Type) (okay bool) {
	for _, n := range faces {
		if n.face == s {
			okay = true
			break
		}
	}
	return
}

func (faces Interfaces) FindMatching(s r.Type) (ret string, err error) {
	var found []string
	for _, n := range faces {
		u, name := n.face, n.name
		if s.Implements(u) || r.PtrTo(s).Implements(u) {
			found = append(found, name)
		}
	}
	if found == nil {
		err = errutil.New("no interface for", s, "in", faces)
	} else {
		ret = strings.Join(found, ",")
	}
	return
}
