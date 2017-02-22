package blocks

import (
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

type ScriptDB interface {
	ReverseCursor(path string) (Cursor, error)
	FindByPath(path string) (interface{}, bool)
}

type Cursor interface {
	HasNext() bool
	GetNext() Location
}

type Location struct {
	Path string
	Data interface{}
}

type DB map[string]interface{}

func (f DB) Add(key []string, data interface{}) {
	path := strings.Join(key, "/")
	f[path] = data
}

func (f DB) FindByPath(path string) (interface{}, bool) {
	x, ok := f[path]
	return x, ok
}

func (f DB) ReverseCursor(path string) (ret Cursor, err error) {
	if len(path) < 0 {
		err = errutil.New("empty path")
	} else {
		c := &DBCurse{f, strings.Split(path, "/"), nil}
		c.advance()
		ret = c
	}
	return
}

type DBCurse struct {
	db    DB
	parts []string
	curr  *Location
}

func (c *DBCurse) HasNext() bool {
	return c.curr != nil
}

func (c *DBCurse) GetNext() Location {
	ret := *c.curr
	if !c.advance() {
		c.curr = nil
	}
	return ret
}

func (c *DBCurse) advance() (okay bool) {
	for len(c.parts) > 0 {
		path := strings.Join(c.parts, "/")
		if data, ok := c.db[path]; ok {
			c.curr = &Location{path, data}
			okay = true
			break
		}
		c.parts = c.parts[:len(c.parts)-1]
	}
	return
}
