# htmlspec
[![Build status](https://travis-ci.org/corbym/htmlspec.svg?branch=master)](https://github.com/corbym/htmlspec)
[![GoDoc](https://godoc.org/github.com/corbym/htmlspec?status.svg)](http://godoc.org/github.com/corbym/htmlspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/htmlspec)](https://goreportcard.com/report/github.com/corbym/htmlspec)

HTML output generator for the BDD framework [GoGiven](https://github.com/corbym/gogiven)
## Import:

```go
import github.com/corbym/htmlspec
```

## Usage:

```go
package foo
import (
	"testing"
	"github.com/corbym/gogiven"
	"github.com/corbym/htmlspec"
	"os"
)

func TestMain(testmain *testing.M) {
	gogiven.Generator = htmlspec.NewTestOutputGenerator()
	runOutput := testmain.Run()
	gogiven.GenerateTestOutput()
	os.Exit(runOutput)
}

... actual tests...

```
