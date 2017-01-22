package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ionous/mars/export"
	"github.com/ionous/mars/export/encode"
	"github.com/ionous/mars/facts"
	"github.com/ionous/mars/std"
	"github.com/ionous/sashimi/util/errutil"
	"log"
	"os"
	"path"
)

func Marshal(src interface{}) (ret string, err error) {
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")
	if e := enc.Encode(src); e != nil {
		err = e
	} else {
		ret = b.String()
	}
	return
}

func write(fileName, m string) {
	log.Println("writing to", fileName)
	if f, e := os.Create(fileName); e != nil {
		log.Println("error", e)
	} else {
		defer f.Close()
		fmt.Fprintln(f, m)
	}
}

// go run libGen.go -base /Users/ionous/Dev/makisu/app/bin
func main() {
	base := flag.String("base", "", "export destination.")
	flag.Parse()

	var mars, exports string
	if *base != "" {
		mars = path.Join(*base, "mars.js")
		exports = path.Join(*base, "exports.js")
	}

	if e := writeMars(mars); e != nil {
		log.Println(e)
	} else if e := writeExports(exports); e != nil {
		log.Println(e)
	}
}

func writeExports(fileName string) (err error) {
	ctx := encode.NewContext()
	if types, e := ctx.AddTypes(export.Export()); e != nil {
		err = errutil.New("types error", e)
	} else {
		if m, e := Marshal(types); e != nil {
			err = errutil.New("marshal error:", e)
		} else {
			if fileName == "" {
				fmt.Println(m)
			} else {
				write(fileName+"on", m)
				write(fileName, "module.exports="+m+";")
			}
		}
	}
	return err
}

func writeMars(fileName string) (err error) {
	ctx := encode.NewContext()
	if libs, e := export.NewLibraries(ctx, export.Export(), std.Std(), facts.Facts()); e != nil {
		err = errutil.New("mars library error", e)
	} else {
		name := path.Base(fileName)
		bundle := export.LibraryBundle{name, libs}
		if data, e := encode.Compute(bundle); e != nil {
			err = errutil.New("compute error", e)
		} else if m, e := Marshal(data); e != nil {
			err = errutil.New("marshal error:", e)
		} else {
			if fileName == "" {
				fmt.Println(m)
			} else {
				write(fileName+"on", m)
				write(fileName, "module.exports="+m+";")
			}
		}
	}
	return err
}
