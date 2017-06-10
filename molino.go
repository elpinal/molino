package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

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
	} else {
		var err error
		arg := os.Args[1]
		source = flag.Arg(0)
		body, err = ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Args = flag.Args()

	molino.Runtime()
	/*
		reader := new(molino.Reader)
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
