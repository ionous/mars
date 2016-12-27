package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ionous/mars/encode"
	"github.com/ionous/mars/std"
	"os"
	"strings"
)

func Marshall(b encode.TypeBlocks) ([]byte, error) {
	return json.MarshalIndent(b, "", " ")
}

func main() {
	b := encode.NewTypeBuilder()
	dst := flag.String("file", "", "export destination.")
	flag.Parse()

	// we get all of the other packages via dependencies
	if e := b.AddPackage(std.Std()); e != nil {
		fmt.Println("error:", e)
	} else {
		tb := b.Build()
		if m, e := Marshall(tb); e != nil {
			fmt.Println("error:", e)
		} else {
			w := os.Stdout
			if *dst != "" {
				fmt.Println("writing to", *dst)
				if f, e := os.Create(*dst); e != nil {
					fmt.Println("error", e)
					return
				} else {
					w = f
					defer f.Close()
				}
			}
			s := "var allTypes =" + strings.Replace(string(m), "Generic", "ObjEval", -1)
			fmt.Fprintln(w, s)
		}
	}
}
