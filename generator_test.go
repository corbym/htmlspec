package htmlspec_test

import (
	"github.com/corbym/gocrest/is"
	. "github.com/corbym/gocrest/then"
	"github.com/corbym/gogiven/base"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/gogiven/testdata"
	"github.com/corbym/htmlspec"
	"testing"
)

var html string
var underTest *htmlspec.TestOutputGenerator

func init() {
	underTest = htmlspec.NewTestOutputGenerator()
}

func TestTestOutputGenerator_Generate(testing *testing.T) {
	fileIsConvertedToHtml()

	AssertThat(testing, html, is.ValueContaining("<title>Generator Test</title>"))
	AssertThat(testing, html, is.ValueContaining("<h1>Generator Test</h1>"))
	AssertThat(testing, html, is.ValueContaining(`<pre class="highlight specification">`))
	AssertThat(testing, html, is.ValueContaining("Given"))
	AssertThat(testing, html, is.ValueContaining("When"))
	AssertThat(testing, html, is.ValueContaining("Then"))
	AssertThat(testing, html, is.ValueContaining("</pre>"))
}

func TestTestOutputGenerator_GenerateConcurrently(testing *testing.T) {
	for i := 0; i < 15; i++ {
		go func() {
			data := newPageData(newSomeMap())

			html := underTest.Generate(data)
			AssertThat(testing, html, is.ValueContaining("<title>Generator Test</title>"))
		}()
	}
}

func TestTestOutputGenerator_FileExtension(t *testing.T) {
	AssertThat(t, underTest.FileExtension(), is.EqualTo(".html"))
}

func TestTestOutputGenerator_Panics(t *testing.T) {
	defer func() {
		recovered := recover()
		AssertThat(t, recovered, is.Not(is.Nil()))
	}()
	data := newPageData(newSomeMap())
	data.SomeMap = nil

	underTest.Generate(data)
}

func fileIsConvertedToHtml() {
	html = underTest.Generate(newPageData(newSomeMap()))
}

func newSomeMap() *base.SomeMap {
	testingT := new(base.TestMetaData)
	return &base.SomeMap{"foo": base.NewSome(testingT,
		"Generator Test",
		base.NewTestMetaData("foo"),
		"GivenWhenThen",
		func(givens testdata.InterestingGivens) {

		}),
	}
}

func newPageData(someMap *base.SomeMap) *generator.PageData {
	return &generator.PageData{
		Title:   "Generator Test",
		SomeMap: someMap,
	}
}
