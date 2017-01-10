package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ionous/mars/facts"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/encode"
	"log"
	"os"
)

func Marshall(src interface{}) (ret string, err error) {
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

// go run libGen.go -file /Users/ionous/Dev/makisu/app/bin/mars.js
func main() {
	dst := flag.String("file", "", "export destination.")
	flag.Parse()

	// we get all of the other packages via dependencies
	if libs, e := encode.AddLibraries(std.Std(), facts.Facts()); e != nil {
		log.Println("error:", e)
	} else if m, e := Marshall(libs); e != nil {
		log.Println("error:", e)
	} else {
		w := os.Stdout
		if *dst != "" {
			log.Println("writing to", *dst)
			if f, e := os.Create(*dst); e != nil {
				log.Println("error", e)
				return
			} else {
				w = f
				defer f.Close()
			}
		}
		s := "var allTypes =" + m
		fmt.Fprintln(w, s)
	}

}
