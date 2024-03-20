package main

import (
	"os"
	"testing"

	"golang.org/x/mod/modfile"
)

func TestPrinter(t *testing.T) {
	badInput, err := os.ReadFile("testdata/bad.mod")
	if err != nil {
		t.Fatal(err)
	}
	goodFile, err := os.ReadFile("testdata/good.mod")
	if err != nil {
		t.Fatal(err)
	}
	got, err := run(badInput)
	if err != nil {
		t.Fatal(err)
	}
	if got != string(goodFile) {
		t.Fatalf("expected %s, got %s", string(goodFile), got)
	}
	if got == string(badInput) {
		t.Fatal("expected change - got none")
	}

}

func run(b []byte) (string, error) {
	mf, err := modfile.Parse("go.mod", b, nil)
	if err != nil {
		return "", err
	}
	p := &printer{}
	p.file(mf)
	return p.String(), nil
}
