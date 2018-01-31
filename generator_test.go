package htmlspec_test

import (
	"github.com/corbym/gocrest/is"
	. "github.com/corbym/gocrest/then"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/htmlspec"
	"testing"
	"bytes"
	"github.com/corbym/gogiven/base"
)

var html string
var underTest *htmlspec.TestOutputGenerator

func init() {
	underTest = htmlspec.NewTestOutputGenerator()
}

func TestTestOutputGenerator_Generate(testing *testing.T) {
	fileIsConvertedToHTML()

	AssertThat(testing, html, is.ValueContaining("<title>Generator Test</title>"))
	AssertThat(testing, html, is.ValueContaining("<h1>Generator Test</h1>"))
	AssertThat(testing, html, is.ValueContaining(`<pre class="highlight specification">`))
	AssertThat(testing, html, is.ValueContaining(`Fooing is best`))
	AssertThat(testing, html, is.ValueContaining(`done with friends`))
	AssertThat(testing, html, is.ValueContaining("given"))
	AssertThat(testing, html, is.ValueContaining("when"))
	AssertThat(testing, html, is.ValueContaining("then"))
	AssertThat(testing, html, is.ValueContaining("</pre>"))
	AssertThat(testing, html, is.ValueContaining(`logkey="foob">`))
	AssertThat(testing, html, is.ValueContaining(`>barb`))
	AssertThat(testing, html, is.ValueContaining(`<th class="key">faff</th>`))
	AssertThat(testing, html, is.ValueContaining(`"interestingGiven">flap`))
}

func TestTestOutputGenerator_GenerateConcurrently(testing *testing.T) {
	data := newPageData(false, false)
	for i := 0; i < 15; i++ {
		go func() {
			buffer := new(bytes.Buffer)
			htmlBytes := underTest.Generate(data)
			buffer.ReadFrom(htmlBytes)

			AssertThat(testing, buffer.String(), is.ValueContaining("<title>Generator Test</title>"))
		}()
	}
}

func TestTestOutputGenerator_FileExtension(t *testing.T) {
	AssertThat(t, underTest.ContentType(), is.EqualTo("text/html"))
}

func fileIsConvertedToHTML() {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(underTest.Generate(newPageData(true, true)))
	html = buffer.String()
}

func newPageData(skipped bool, failed bool) generator.PageData {
	testData := make(map[string]generator.TestData)
	capturedIO := make(map[interface{}]interface{})
	capturedIO["foob"] = "barb"
	interestingGivens := make(map[interface{}]interface{})
	interestingGivens["faff"] = "flap"
	parsedContent := base.ParsedTestContent{
		GivenWhenThen: []string{"given", "when", "then"},
		Comment:       []string{"Fooing is best", "done with friends"},
	}
	testData["test title"] = generator.TestData{
		TestTitle:         "test title",
		ParsedTestContent: parsedContent,
		CapturedIO:        capturedIO,
		InterestingGivens: interestingGivens,
		TestResult: generator.TestResult{
			Failed:     failed,
			Skipped:    skipped,
			TestOutput: "well arighty then",
			TestID:     "abc2124",
		},
	}
	return generator.PageData{
		TestResults: testData,
		Title:       "Generator Test",
	}
}
