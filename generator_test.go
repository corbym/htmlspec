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

	AssertThat(testing, html, is.ValueContaining("<title>Head Generator Test Title</title>"))
	AssertThat(testing, html, is.ValueContaining("<h1>Head Generator Test Title</h1>"))
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
			buffer := generateData(data)

			AssertThat(testing, buffer.String(), is.ValueContaining("<title>Head Generator Test Title</title>"))
		}()
	}
}

func TestTestOutputGenerator_FileExtension(t *testing.T) {
	AssertThat(t, underTest.ContentType(), is.EqualTo("text/html"))
}

func TestTestOutputGenerator_GenerateIndex(t *testing.T) {
	t.Skipf("WIP")
	testData := make([] generator.TestData, 2)
	testData = append(testData, newTestData("First", "abc2124", true, false))
	testData = append(testData, newTestData("Second", "abc2443", true, false))

	someIndexData := []generator.IndexData{
		{Title: "Wombat Test", TestFileName: "/bar/baz/wombat_test.go", TestData: testData},
		{Title: "Normal Bat Test", TestFileName: "/bar/baz/bat_test.go", TestData: testData},
	}
	generatedIndex := generateIndexData(someIndexData)

	AssertThat(t, generatedIndex, is.ValueContaining("<title>Package Name</title>"))
	AssertThat(t, generatedIndex, is.ValueContaining(`<a href="./wombat_test.html" class="">Wombat Test`))
	AssertThat(t, generatedIndex, is.ValueContaining(`<a href="./wombat_bat_test.html#abc2124" class="">First`))
	AssertThat(t, generatedIndex, is.ValueContaining(`<a href="./wombat_bat_test.html#abc2443" class="">Second`))

	AssertThat(t, generatedIndex, is.ValueContaining(`<a href="./wombat_test.html">Normal Bat Test`))
	AssertThat(t, generatedIndex, is.ValueContaining(`<a href="./bat_test.html#abc2124" class="">First`))
	AssertThat(t, generatedIndex, is.ValueContaining(`<a href="./bat_test.html#abc2443" class="">Second`))
}

func generateData(data generator.PageData) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	htmlBytes := underTest.Generate(data)
	buffer.ReadFrom(htmlBytes)
	return buffer
}

func generateIndexData(data []generator.IndexData) string {
	buffer := new(bytes.Buffer)
	htmlBytes := underTest.GenerateIndex(data)
	buffer.ReadFrom(htmlBytes)
	return buffer.String()
}

func fileIsConvertedToHTML() {
	html = generateData(newPageData(true, true)).String()
}

func newPageData(skipped bool, failed bool) generator.PageData {
	testData := make([]generator.TestData, 1)
	testData = append(testData, newTestData("Test Title", "abc2124", failed, skipped))
	return generator.PageData{
		TestData: testData,
		Title:    "Head Generator Test Title",
	}
}

func newTestData(testTitle string, testId string, failed bool, skipped bool) generator.TestData {
	capturedIO := make(map[interface{}]interface{})
	capturedIO["foob"] = "barb"
	interestingGivens := make(map[interface{}]interface{})
	interestingGivens["faff"] = "flap"
	parsedContent := base.ParsedTestContent{
		GivenWhenThen: []string{"given", "when", "then"},
		Comment:       []string{"Fooing is best", "done with friends"},
	}
	testData := generator.TestData{
		TestTitle:         testTitle,
		ParsedTestContent: parsedContent,
		CapturedIO:        capturedIO,
		InterestingGivens: interestingGivens,
		TestResult: generator.TestResult{
			Failed:     failed,
			Skipped:    skipped,
			TestOutput: "well arighty then",
			TestID:     testId,
		},
	}
	return testData
}
