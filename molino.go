package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elpinal/molino/src/go/molino"
)

var e = flag.String("e", "", "One line of program")

func main() {
	flag.Parse()
	var body []byte
	var source string
	if *e != "" {
		body = []byte(*e)
		source = "argument"
	} else if flag.NArg() != 1 {
		log.Fatal("1 argument required")
	} else {
		var err error
		source = flag.Arg(0)
		body, err = ioutil.ReadFile(source)
		if err != nil {
			log.Fatal(err)
		}
	}

	molino.Runtime()
	reader := new(molino.Reader)
	reader.Init(string(body))
	var ret interface{}
	for r, eof, err := reader.Read(); !eof; r, eof, err = reader.Read() {
		if err != nil {
			// log.SetFlags(log.Lshortfile)
			log.Fatal(err)
		}
		fmt.Println(r)
		ret = molino.Eval(r)
		fmt.Println(ret)
	}
}
