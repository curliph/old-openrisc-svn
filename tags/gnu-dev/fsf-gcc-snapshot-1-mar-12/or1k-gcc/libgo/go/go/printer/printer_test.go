// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package printer

import (
	"bytes"
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"
)

const (
	dataDir  = "testdata"
	tabwidth = 8
)

var update = flag.Bool("update", false, "update golden files")

var fset = token.NewFileSet()

func lineString(text []byte, i int) string {
	i0 := i
	for i < len(text) && text[i] != '\n' {
		i++
	}
	return string(text[i0:i])
}

type checkMode uint

const (
	export checkMode = 1 << iota
	rawFormat
)

func runcheck(t *testing.T, source, golden string, mode checkMode) {
	// parse source
	prog, err := parser.ParseFile(fset, source, nil, parser.ParseComments)
	if err != nil {
		t.Error(err)
		return
	}

	// filter exports if necessary
	if mode&export != 0 {
		ast.FileExports(prog) // ignore result
		prog.Comments = nil   // don't print comments that are not in AST
	}

	// determine printer configuration
	cfg := Config{Tabwidth: tabwidth}
	if mode&rawFormat != 0 {
		cfg.Mode |= RawFormat
	}

	// format source
	var buf bytes.Buffer
	if err := cfg.Fprint(&buf, fset, prog); err != nil {
		t.Error(err)
	}
	res := buf.Bytes()

	// update golden files if necessary
	if *update {
		if err := ioutil.WriteFile(golden, res, 0644); err != nil {
			t.Error(err)
		}
		return
	}

	// get golden
	gld, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Error(err)
		return
	}

	// compare lengths
	if len(res) != len(gld) {
		t.Errorf("len = %d, expected %d (= len(%s))", len(res), len(gld), golden)
	}

	// compare contents
	for i, line, offs := 0, 1, 0; i < len(res) && i < len(gld); i++ {
		ch := res[i]
		if ch != gld[i] {
			t.Errorf("%s:%d:%d: %s", source, line, i-offs+1, lineString(res, offs))
			t.Errorf("%s:%d:%d: %s", golden, line, i-offs+1, lineString(gld, offs))
			t.Error()
			return
		}
		if ch == '\n' {
			line++
			offs = i + 1
		}
	}
}

func check(t *testing.T, source, golden string, mode checkMode) {
	// start a timer to produce a time-out signal
	tc := make(chan int)
	go func() {
		time.Sleep(10 * time.Second) // plenty of a safety margin, even for very slow machines
		tc <- 0
	}()

	// run the test
	cc := make(chan int)
	go func() {
		runcheck(t, source, golden, mode)
		cc <- 0
	}()

	// wait for the first finisher
	select {
	case <-tc:
		// test running past time out
		t.Errorf("%s: running too slowly", source)
	case <-cc:
		// test finished within alloted time margin
	}
}

type entry struct {
	source, golden string
	mode           checkMode
}

// Use gotest -update to create/update the respective golden files.
var data = []entry{
	{"empty.input", "empty.golden", 0},
	{"comments.input", "comments.golden", 0},
	{"comments.input", "comments.x", export},
	{"linebreaks.input", "linebreaks.golden", 0},
	{"expressions.input", "expressions.golden", 0},
	{"expressions.input", "expressions.raw", rawFormat},
	{"declarations.input", "declarations.golden", 0},
	{"statements.input", "statements.golden", 0},
	{"slow.input", "slow.golden", 0},
}

func TestFiles(t *testing.T) {
	for i, e := range data {
		source := filepath.Join(dataDir, e.source)
		golden := filepath.Join(dataDir, e.golden)
		check(t, source, golden, e.mode)
		// TODO(gri) check that golden is idempotent
		//check(t, golden, golden, e.mode)
		if testing.Short() && i >= 3 {
			break
		}
	}
}

// TestLineComments, using a simple test case, checks that consequtive line
// comments are properly terminated with a newline even if the AST position
// information is incorrect.
//
func TestLineComments(t *testing.T) {
	const src = `// comment 1
	// comment 2
	// comment 3
	package main
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err) // error in test
	}

	var buf bytes.Buffer
	fset = token.NewFileSet() // use the wrong file set
	Fprint(&buf, fset, f)

	nlines := 0
	for _, ch := range buf.Bytes() {
		if ch == '\n' {
			nlines++
		}
	}

	const expected = 3
	if nlines < expected {
		t.Errorf("got %d, expected %d\n", nlines, expected)
		t.Errorf("result:\n%s", buf.Bytes())
	}
}

// Verify that the printer can be invoked during initialization.
func init() {
	const name = "foobar"
	var buf bytes.Buffer
	if err := Fprint(&buf, fset, &ast.Ident{Name: name}); err != nil {
		panic(err) // error in test
	}
	// in debug mode, the result contains additional information;
	// ignore it
	if s := buf.String(); !debug && s != name {
		panic("got " + s + ", want " + name)
	}
}

// Verify that the printer doesn't crash if the AST contains BadXXX nodes.
func TestBadNodes(t *testing.T) {
	const src = "package p\n("
	const res = "package p\nBadDecl\n"
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err == nil {
		t.Error("expected illegal program") // error in test
	}
	var buf bytes.Buffer
	Fprint(&buf, fset, f)
	if buf.String() != res {
		t.Errorf("got %q, expected %q", buf.String(), res)
	}
}

// Print and parse f with 
func testComment(t *testing.T, f *ast.File, srclen int, comment *ast.Comment) {
	f.Comments[0].List[0] = comment
	var buf bytes.Buffer
	for offs := 0; offs <= srclen; offs++ {
		buf.Reset()
		// Printing f should result in a correct program no
		// matter what the (incorrect) comment position is.
		if err := Fprint(&buf, fset, f); err != nil {
			t.Error(err)
		}
		if _, err := parser.ParseFile(fset, "", buf.Bytes(), 0); err != nil {
			t.Fatalf("incorrect program for pos = %d:\n%s", comment.Slash, buf.String())
		}
		// Position information is just an offset.
		// Move comment one byte down in the source.
		comment.Slash++
	}
}

// Verify that the printer produces always produces a correct program
// even if the position information of comments introducing newlines
// is incorrect.
func TestBadComments(t *testing.T) {
	const src = `
// first comment - text and position changed by test
package p
import "fmt"
const pi = 3.14 // rough circle
var (
	x, y, z int = 1, 2, 3
	u, v float64
)
func fibo(n int) {
	if n < 2 {
		return n /* seed values */
	}
	return fibo(n-1) + fibo(n-2)
}
`

	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		t.Error(err) // error in test
	}

	comment := f.Comments[0].List[0]
	pos := comment.Pos()
	if fset.Position(pos).Offset != 1 {
		t.Error("expected offset 1") // error in test
	}

	testComment(t, f, len(src), &ast.Comment{pos, "//-style comment"})
	testComment(t, f, len(src), &ast.Comment{pos, "/*-style comment */"})
	testComment(t, f, len(src), &ast.Comment{pos, "/*-style \n comment */"})
	testComment(t, f, len(src), &ast.Comment{pos, "/*-style comment \n\n\n */"})
}
