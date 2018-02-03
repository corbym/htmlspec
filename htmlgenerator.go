package htmlspec

import (
	"bytes"
	"github.com/corbym/gogiven/generator"
	"html/template"
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

var lastError error
//NewTestOutputGenerator creates a template that is used to generate the html output.
func NewTestOutputGenerator() *TestOutputGenerator {
	outputGenerator := new(TestOutputGenerator)
	template := template.New(baseTemplateName)
	template.Parse(safeStringConverter(Asset("resources/htmltemplate.gtl")))
	template.Parse(safeStringConverter(Asset("resources/interestinggivens.gtl")))
	template.Parse(safeStringConverter(Asset("resources/capturedio.gtl")))
	template.Parse(safeStringConverter(Asset("resources/style.gtl")))
	template.Parse(safeStringConverter(Asset("resources/test-body.gtl")))
	template.Parse(safeStringConverter(Asset("resources/contents.gtl")))
	template.Parse(safeStringConverter(Asset("resources/javascript.gtl")))

	if lastError != nil {
		panic(lastError.Error())
	}
	outputGenerator.template = template
	return outputGenerator
}
func safeStringConverter(asset []byte, err error) string {
	if err != nil {
		lastError = err
	}
	return string(asset[:])
}

// ContentType for the output generated.
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
