package htmlspec

import (
	"bytes"
	"github.com/corbym/gogiven/generator"
	"html/template"
	"os"
	"path/filepath"
	"io"
)

const goPathEnvKey = "GOPATH"
const resourcesPath = "src/github.com/corbym/htmlspec/resources/"
const baseTemplateName = "base"

//TestOutputGenerator is an implementation of the GoGivensOutputGenerator that generates an html file per
// test. It is thread safe between goroutines.
type TestOutputGenerator struct {
	generator.GoGivensOutputGenerator
	template *template.Template
}

//NewTestOutputGenerator creates a template that is used to generate the html output.
func NewTestOutputGenerator() *TestOutputGenerator {
	goPath := os.Getenv(goPathEnvKey)
	outputGenerator := new(TestOutputGenerator)
	outputGenerator.template = template.Must(template.ParseFiles(
		filepath.Join(goPath, resourcesPath+"htmltemplate.gtl"),
		filepath.Join(goPath, resourcesPath+"interestinggivens.gtl"),
		filepath.Join(goPath, resourcesPath+"capturedio.gtl"),
		filepath.Join(goPath, resourcesPath+"style.gtl"),
		filepath.Join(goPath, resourcesPath+"test-body.gtl"),
		filepath.Join(goPath, resourcesPath+"contents.gtl"),
		filepath.Join(goPath, resourcesPath+"javascript.gtl"),
	))
	return outputGenerator
}

// FileExtension for the output generated.
func (outputGenerator *TestOutputGenerator) ContentType() string {
	return "text/html"
}

// Generate generates html output for a test. The return string contains the html
// that goes into the output file generated in gogivens.GenerateTestOutput().
// The function panics if the template cannot be generated.
func (outputGenerator *TestOutputGenerator) Generate(pageData generator.PageData) io.Reader {
	var buffer = new(bytes.Buffer)
	outputGenerator.template.ExecuteTemplate(buffer, baseTemplateName, pageData)
	return buffer
}
