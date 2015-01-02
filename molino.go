package main

import (
	"./src/go/molino/lang"
	"flag"
	"io/ioutil"
	"log"
	"os"
//	"fmt"
)

var fs = flag.NewFlagSet(os.Args[0], 1)
var e = fs.String("e", "", "One line of program")

func main() {
	fs.Parse(os.Args[1:])
	var body []byte
	var source string
	if *e != "" {
		body = []byte(*e)
		source = "argument"
	} else {
		var err error
		arg := os.Args[1]
		source = fs.Arg(0)
		body, err = ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Args = fs.Args()

	lang.Runtime()
	/*
	reader := new(lang.Reader)
	reader.Init(string(body))
	var ret interface{}
	for r, eof, err := reader.Read(); !eof; r, eof, err = reader.Read() {
		if err != nil {
			// log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}
		fmt.Println(r)
		ret = eval(r)
	}
*/
	_, _ = source, body
}
