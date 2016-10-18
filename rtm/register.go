package rtm

import "reflect"

// RegisterName matches gob
type RegisterName func(name string, value interface{})

func RegisterTypes(reg RegisterName, libs ...interface{}) {
	for _, lib := range libs {
		t := reflect.TypeOf(lib)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			v := reflect.New(f.Type.Elem())
			reg(f.Name, v.Elem().Interface())
		}
	}
}
