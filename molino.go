package main

import (
	"./src/go/molino/lang"
	"flag"
	"io/ioutil"
	"log"
	"os"
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
}
