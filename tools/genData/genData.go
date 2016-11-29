package main

import (
	"encoding/json"
	"fmt"
	"github.com/ionous/mars/encode"
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	r "reflect"
)

func Marshall(cmd encode.DataBlock) ([]byte, error) {
	return json.MarshalIndent(cmd, "", " ")
}

func main() {
	var d = script.Script{
		Name: "MyScript",
		Statements: []backend.Spec{
			//script.The("kinds", script.Can("test").And("testing").RequiresTwo("things")),
			script.Understand("look|l at {{something}}").As("test"),
		},
	}
	if cmd, e := encode.ComputeCmd(r.ValueOf(d)); e != nil {
		fmt.Println("error", e)
	} else if m, e := Marshall(cmd); e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println(string(m))
	}
}

//  "name": "Fragments"
//     "phrase": "[fragment]"
//     "uses": "" <-- blank , should be "Fragment"
//     // "usesArray": true <-=- should be true!
