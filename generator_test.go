package htmlspec_test

import (
	"github.com/corbym/gocrest/is"
	. "github.com/corbym/gocrest/then"
	"testing"
	"github.com/corbym/gogiven/base"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/gogiven/testdata"
	"github.com/corbym/htmlspec"
)

var rawContent string
var html string

func TestGeneratorCreatesTestFirstTestHtml(testing *testing.T) {
	fileIsConvertedToHtml()

	AssertThat(testing, html, is.ValueContaining("<title>Generator Test</title>"))
	AssertThat(testing, html, is.ValueContaining("<h1>Generator Test</h1>"))
	AssertThat(testing, html, is.ValueContaining("Given"))
	AssertThat(testing, html, is.ValueContaining("When"))
	AssertThat(testing, html, is.ValueContaining("Then"))
}

func fileIsConvertedToHtml() {
	testingT := new(base.TestMetaData)
	someMap := &base.SomeMap{"foo": base.NewSome(testingT,
		"Generator Test",
		base.NewTestMetaData("foo"),
		"Given\nWhen\nThen",
		func(givens *testdata.InterestingGivens) {

		}),
	}
	gen := new(htmlspec.TestOutputGenerator)
	data := &generator.PageData{
		Title:   "Generator Test",
		SomeMap: someMap,
	}
	html = gen.Generate(data)
}
