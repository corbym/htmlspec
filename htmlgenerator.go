package htmlspec

import (
	"strings"
	"bytes"
	"github.com/corbym/gogiven/generator"
	"html/template"
	"io"
	"sync"
)

const goPathEnvKey = "GOPATH"
const resourcesPath = "src/github.com/corbym/htmlspec/resources/"
const baseTemplateName = "base"

//TestOutputGenerator is an implementation of the GoGivensOutputGenerator that generates an html file per
// test. It is thread safe between goroutines.
type TestOutputGenerator struct {
	sync.RWMutex
	generator.GoGivensOutputGenerator
	template       *template.Template
	generatedPages map[string]generator.PageData
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
	outputGenerator.generatedPages = make(map[string]generator.PageData, 10)
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
	outputGenerator.Lock()
	defer func() { outputGenerator.Unlock() }()

	outputGenerator.generatedPages[pageData.Title] = pageData //spike to see if this works
	return buffer
}

//GenerateIndex generates an index of all the data from the tests in HTML format.
func (outputGenerator *TestOutputGenerator) GenerateIndex() io.Reader {
	var content string
	outputGenerator.RLock()
	defer func() { outputGenerator.RUnlock() }()
	for _, page := range outputGenerator.generatedPages {
		content += page.Title + "<BR>"
	}
	return strings.NewReader(content)
}
