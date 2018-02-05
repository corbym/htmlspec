package htmlspec

import (
	"strings"
	"bytes"
	"github.com/corbym/gogiven/generator"
	"html/template"
	"io"
)

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
	htmlTemplate := template.New(baseTemplateName)
	htmlTemplate.Parse(safeStringConverter(Asset("resources/htmltemplate.gtl")))
	htmlTemplate.Parse(safeStringConverter(Asset("resources/interestinggivens.gtl")))
	htmlTemplate.Parse(safeStringConverter(Asset("resources/capturedio.gtl")))
	htmlTemplate.Parse(safeStringConverter(Asset("resources/style.gtl")))
	htmlTemplate.Parse(safeStringConverter(Asset("resources/test-body.gtl")))
	htmlTemplate.Parse(safeStringConverter(Asset("resources/contents.gtl")))
	htmlTemplate.Parse(safeStringConverter(Asset("resources/javascript.gtl")))

	if lastError != nil {
		panic(lastError.Error())
	}
	outputGenerator.template = htmlTemplate
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
// that goes into the output generated in gogivens.GenerateTestOutput().
// The function panics if the template cannot be generated.
func (outputGenerator *TestOutputGenerator) Generate(pageData generator.PageData) (io.Reader) {
	var buffer = new(bytes.Buffer)
	outputGenerator.template.ExecuteTemplate(buffer, baseTemplateName, pageData)
	return buffer
}

//GenerateIndex generates an index of all the data from the tests in HTML format.
func (outputGenerator *TestOutputGenerator) GenerateIndex(indexData []generator.IndexData) io.Reader {
	var content string
	for _, idx := range indexData {
		content += idx.Title + "<BR>"
	}
	return strings.NewReader(content)
}
