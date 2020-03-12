package main

/*

This file is adapted from golang.org/x/mod/modfile/printer.go
but implemented by printing

*/

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
)

// A printer collects the state during printing of a file or expression.
type printer struct {
	bytes.Buffer                   // output buffer
	comment      []modfile.Comment // pending end-of-line comments
	margin       int               // left margin (indent), a number of tabs
}

// printf prints to the buffer.
func (p *printer) printf(format string, args ...interface{}) {
	fmt.Fprintf(p, format, args...)
}

// indent returns the position on the current line, in bytes, 0-indexed.
func (p *printer) indent() int {
	b := p.Bytes()
	n := 0
	for n < len(b) && b[len(b)-1-n] != '\n' {
		n++
	}
	return n
}

// newline ends the current line, flushing end-of-line comments.
func (p *printer) newline() {
	if len(p.comment) > 0 {
		p.printf(" ")
		for i, com := range p.comment {
			if i > 0 {
				p.trim()
				p.printf("\n")
				for i := 0; i < p.margin; i++ {
					p.printf("\t")
				}
			}
			p.printf("%s", strings.TrimSpace(com.Token))
		}
		p.comment = p.comment[:0]
	}

	p.trim()
	p.printf("\n")
	for i := 0; i < p.margin; i++ {
		p.printf("\t")
	}
}

// trim removes trailing spaces and tabs from the current line.
func (p *printer) trim() {
	// Remove trailing spaces and tabs from line we're about to end.
	b := p.Bytes()
	n := len(b)
	for n > 0 && (b[n-1] == '\t' || b[n-1] == ' ') {
		n--
	}
	p.Truncate(n)
}

// file formats the given file into the print buffer.
func (p *printer) file(f *modfile.File) {
	sort.Slice(f.Require, func(i, j int) bool { return f.Require[i].Mod.Path < f.Require[j].Mod.Path })
	sort.Slice(f.Exclude, func(i, j int) bool { return f.Exclude[i].Mod.Path < f.Exclude[j].Mod.Path })
	sort.Slice(f.Replace, func(i, j int) bool { return f.Replace[i].Old.Path < f.Replace[j].Old.Path })

	for _, com := range f.Syntax.Before {
		p.printf("%s", strings.TrimSpace(com.Token))
		p.newline()
	}

	p.expr(f.Module.Syntax)
	p.newline()
	p.newline()
	p.expr(f.Go.Syntax)
	p.newline()

	if len(f.Require) > 1 {
		p.newline()
		p.printf("require (")
		p.margin++
		for _, r := range f.Require {
			if !r.Syntax.InBlock {
				r.Syntax.Token = r.Syntax.Token[1:len(r.Syntax.Token)]
			}
			p.newline()
			p.expr(r.Syntax)
		}
		p.margin--
		p.newline()
		p.printf(")")
		p.newline()
	} else if len(f.Require) == 1 {
		r := f.Require[0]
		if r.Syntax.InBlock {
			r.Syntax.Token = append([]string{"require"}, r.Syntax.Token...)
		}
		p.newline()
		p.expr(r.Syntax)
		p.newline()
	}
	if len(f.Exclude) > 1 {
		p.newline()
		p.printf("exclude (")
		p.margin++
		for _, r := range f.Exclude {
			p.newline()
			p.expr(r.Syntax)
		}
		p.margin--
		p.newline()
		p.printf(")")
		p.newline()
	} else if len(f.Exclude) == 1 {
		r := f.Exclude[0]
		if r.Syntax.InBlock {
			r.Syntax.Token = append([]string{"exclude"}, r.Syntax.Token...)
		}
		p.newline()
		p.expr(r.Syntax)
		p.newline()
	}
	if len(f.Replace) > 1 {
		p.newline()
		p.printf("replace (")
		p.margin++
		for _, r := range f.Replace {
			if !r.Syntax.InBlock {
				r.Syntax.Token = r.Syntax.Token[1:len(r.Syntax.Token)]
			}
			p.newline()
			p.expr(r.Syntax)
		}
		p.margin--
		p.newline()
		p.printf(")")
		p.newline()
	} else if len(f.Replace) == 1 {
		r := f.Replace[0]
		if r.Syntax.InBlock {
			r.Syntax.Token = append([]string{"replace"}, r.Syntax.Token...)
		}
		p.newline()
		p.expr(r.Syntax)
		p.newline()
	}

}

func (p *printer) expr(x modfile.Expr) {
	// Emit line-comments preceding this expression.
	if before := x.Comment().Before; len(before) > 0 {
		// Want to print a line comment.
		// Line comments must be at the current margin.
		p.trim()
		if p.indent() > 0 {
			// There's other text on the line. Start a new line.
			p.printf("\n")
		}
		// Re-indent to margin.
		for i := 0; i < p.margin; i++ {
			p.printf("\t")
		}
		for _, com := range before {
			p.printf("%s", strings.TrimSpace(com.Token))
			p.newline()
		}
	}

	switch x := x.(type) {
	default:
		panic(fmt.Errorf("printer: unexpected type %T", x))

	case *modfile.CommentBlock:
		// done

	case *modfile.LParen:
		p.printf("(")
	case *modfile.RParen:
		p.printf(")")

	case *modfile.Line:
		sep := ""
		for _, tok := range x.Token {
			p.printf("%s%s", sep, tok)
			sep = " "
		}

	case *modfile.LineBlock:
		for _, tok := range x.Token {
			p.printf("%s ", tok)
		}
		p.expr(&x.LParen)
		p.margin++
		for _, l := range x.Line {
			p.newline()
			p.expr(l)
		}
		p.margin--
		p.newline()
		p.expr(&x.RParen)
	}

	// Queue end-of-line comments for printing when we
	// reach the end of the line.
	p.comment = append(p.comment, x.Comment().Suffix...)
}
