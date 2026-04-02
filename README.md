# htmlspec
[![CI](https://github.com/corbym/htmlspec/actions/workflows/ci.yml/badge.svg)](https://github.com/corbym/htmlspec/actions/workflows/ci.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/corbym/htmlspec.svg)](https://pkg.go.dev/github.com/corbym/htmlspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/htmlspec)](https://goreportcard.com/report/github.com/corbym/htmlspec)

HTML output generator for the BDD framework [GoGiven](https://github.com/corbym/gogiven).

Generates a responsive HTML5 report per test file, plus an index page, with pass/fail/skip status badges and collapsible sections for interesting givens and captured I/O.

## Requirements

Go 1.21+

## Install:

```
go get github.com/corbym/htmlspec
```

## Usage:

```go
package foo

import (
	"os"
	"testing"

	"github.com/corbym/gogiven"
	"github.com/corbym/htmlspec"
)

func TestMain(testmain *testing.M) {
	gogiven.Generator = htmlspec.NewHTMLOutputGenerator()
	runOutput := testmain.Run()
	gogiven.GenerateTestOutput()
	os.Exit(runOutput)
}

// ... actual tests ...
```
