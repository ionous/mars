

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ionous/mars/encode"
	"github.com/ionous/mars/facts"
	"github.com/ionous/mars/std"
	"os"
)

func Marshall(src encode.TypeBlocks) (ret string, err error) {
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

func main() {
		
	libs := encode.NewLibEncoder()
	dst := flag.String("file", "", "export destination.")
	flag.Parse()

	// we get all of the other packages via dependencies
	if e := libs.AddPackage(std.Std()); e != nil {
		fmt.Println("error:", e)
	} else if e := libs.AddPackage(facts.Facts()); e != nil {
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
			s := "var allTypes =" + m
			fmt.Fprintln(w, s)
		}
	}
}
