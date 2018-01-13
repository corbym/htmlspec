package htmlspec

import (
	"bytes"
	"github.com/corbym/gogiven/generator"
	"html/template"
	"os"
	"path/filepath"
)

const goPathEnvKey = "GOPATH"
const resourcesPath = "src/github.com/corbym/htmlspec/resources/"
const baseTemplateName = "base"

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
	goPath := os.Getenv(goPathEnvKey)
	tmpl := template.Must(template.ParseFiles(
		filepath.Join(goPath, resourcesPath+"htmltemplate.gtl"),
		filepath.Join(goPath, resourcesPath+"capturedio.gtl"),
		filepath.Join(goPath, resourcesPath+"interestinggivens.gtl"),
		filepath.Join(goPath, resourcesPath+"style.gtl"),
		filepath.Join(goPath, resourcesPath+"test-body.gtl"),
		filepath.Join(goPath, resourcesPath+"contents.gtl"),
		filepath.Join(goPath, resourcesPath+"javascript.gtl"),
	))
	var buffer = new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buffer, baseTemplateName, pageData)
	if err != nil {
		panic(err.Error())
	}
	return buffer.String()
}
