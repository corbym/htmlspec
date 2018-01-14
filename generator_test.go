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

var rawContent string
var html string
var gen = htmlspec.NewTestOutputGenerator()

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

func TestTestOutputGenerator_FileExtension(t *testing.T) {
	AssertThat(t, gen.FileExtension(), is.EqualTo(".html"))
}
func TestTestOutputGenerator_Panics(t *testing.T) {
	defer func() {
		recovered := recover()
		AssertThat(t, recovered, is.Not(is.Nil()))
	}()
	data.SomeMap = nil
	fileIsConvertedToHtml()
}

var testingT = new(base.TestMetaData)
var someMap = &base.SomeMap{"foo": base.NewSome(testingT,
	"Generator Test",
	base.NewTestMetaData("foo"),
	"GivenWhenThen",
	func(givens testdata.InterestingGivens) {

	}),
}

var data = &generator.PageData{
	Title:   "Generator Test",
	SomeMap: someMap,
}

func fileIsConvertedToHtml() {
	html = gen.Generate(data)
}
