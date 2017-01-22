package encode

import (
	"github.com/ionous/mars/script/backend"
)

func RecodeScript(scripts []backend.Declaration) (ret []DataBlock, err error) {
	for _, script := range scripts {
		if data, e := Compute(script); e != nil {
			err = e
			break
		} else {
			ret = append(ret, data)
		}
	}
	return
}
