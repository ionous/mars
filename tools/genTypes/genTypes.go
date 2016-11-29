package main

import (
	"encoding/json"
	"fmt"
	"github.com/ionous/mars/encode"
	"github.com/ionous/mars/std"
)

func Marshall(b encode.TypeBlocks) ([]byte, error) {
	return json.MarshalIndent(b, "", " ")
}

func main() {
	//story := flag.String("story", "", "select the story to play.")
	// verbose := flag.Bool("verbose", false, "prints log output when true.")
	// 	text := flag.Bool("text", false, "uses the simpler text console when true.")
	// 	dump := flag.Bool("dump", false, "dump the model.")
	// 	load := flag.Bool("load", false, "load the story save game.")
	// 	flag.Parse()
	// flag.PrintDefaults()
	b := encode.NewTypeBuilder()

	// we get all of the other packages via dependencies
	if e := b.AddPackage(std.Std()); e != nil {
		fmt.Println("error:", e)
	} else {
		tb := b.Build()
		if m, e := Marshall(tb); e != nil {
			fmt.Println("error:", e)
		} else {
			fmt.Println(string(m))
		}
	}
}
