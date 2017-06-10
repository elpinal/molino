package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elpinal/molino/src/go/molino"
)

var e = flag.String("e", "", "One line of program")

func main() {
	flag.Parse()
	body := *e
	if body == "" {
		if flag.NArg() != 1 {
			fmt.Fprintln(os.Stderr, "1 argument required")
			os.Exit(1)
		}
		source := flag.Arg(0)
		b, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		body = string(b)
	}

	molino.Runtime()
	reader := molino.Reader{}
	reader.Init(body)
	var ret interface{}
	for r, eof, err := reader.Read(); !eof; r, eof, err = reader.Read() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(r)
		ret = molino.Eval(r)
		fmt.Println(ret)
	}
}
