package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	f := flag.String("f", "", "filename")
	w := flag.Bool("w", false, "overwrite source file")
	flag.Parse()
	b, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Fatal(err)
	}
	mf, err := modfile.Parse(*f, b, nil)
	if err != nil {
		log.Fatal(err)
	}

	p := &printer{}
	p.file(mf)
	if *w {
		of, err := os.Create(*f)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = of.Write(p.Bytes()); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(p.String())
	}
}
