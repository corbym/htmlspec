package htmlspec

import (
	"path/filepath"
	"bytes"
	"os"
	"html/template"
	"github.com/corbym/gogiven/generator"
)

//TestOutputGenerator is an implmentation of the GoGivensOutputGenerator that generates an html file per
// test.
type TestOutputGenerator struct {
	generator.GoGivensOutputGenerator
}

// FileExtension for the output generated.

func (generator *TestOutputGenerator) FileExtension() string {
	return ".html"
}

// Generate generates the default output for a test. The return string contains the html
// that goes into the output file generated in gogivens.GenerateTestOutput()
func (generator *TestOutputGenerator) Generate(pageData *generator.PageData) string {
	goPath := os.Getenv("GOPATH")
	tmpl := template.Must(template.ParseFiles(
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/htmltemplate.gtl"),
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/capturedio.gtl"),
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/interestinggivens.gtl"),
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/style.gtl"),
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/test-body.gtl"),
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/contents.gtl"),
		filepath.Join(goPath, "src/github.com/corbym/htmlspec/resources/javascript.gtl"),
	))
	var buffer = new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buffer, "base", pageData)
	if err != nil {
		panic(err.Error())
	}
	return buffer.String()
}
