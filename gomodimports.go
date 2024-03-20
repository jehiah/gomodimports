package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	f := flag.String("f", "", "path to go.mod")
	w := flag.Bool("w", false, "overwrite source file")
	l := flag.Bool("l", false, "list files whose formatting differs")
	flag.Parse()
	b, err := os.ReadFile(*f)
	if err != nil {
		log.Fatal(err)
	}
	mf, err := modfile.Parse(*f, b, nil)
	if err != nil {
		log.Fatal(err)
	}

	p := &printer{}
	p.file(mf)
	if *l {
		if !bytes.Equal(p.Bytes(), b) {
			fmt.Println(*f)
		}
	} else if *w {
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
