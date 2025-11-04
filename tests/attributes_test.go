package tests

import (
	"testing"

	. "github.com/delaneyj/gostar/elements"
	"github.com/stretchr/testify/assert"
)

func TestBoolAttributes(t *testing.T) {
	run(t, []result{
		{
			Expected: "<input disabled>",
			Actual:   INPUT().DISABLED(),
		},
		{
			Expected: "<button></button>",
			Actual:   BUTTON().IfDISABLED(false),
		},
		{
			Expected: "<select multiple></select>",
			Actual:   SELECT().MULTIPLE(),
		},
		{
			Expected: "<option selected></option>",
			Actual:   OPTION().SELECTED(),
		},
		{
			Expected: "<textarea readonly></textarea>",
			Actual:   TEXTAREA().READONLY(),
		},
		{
			Expected: "<form novalidate></form>",
			Actual:   FORM().NOVALIDATE(),
		},
		{
			Expected: "<iframe allowfullscreen></iframe>",
			Actual:   IFRAME().ALLOWFULLSCREEN(),
		},
		{
			Expected: "<fieldset disabled></fieldset>",
			Actual:   FIELDSET().DISABLED(),
		},
		{
			Expected: "<form disabled></form>",
			Actual:   FORM().BoolAttr("disabled"),
		},
		{
			Expected: "<input>",
			Actual:   INPUT().IfBoolAttr(false, "disabled"),
		},
		{
			Expected: "<video autoplay muted></video>",
			Actual:   VIDEO().AUTOPLAY().MUTED(),
		},
		{
			Expected: "<video autoplay muted></video>",
			Actual:   VIDEO().AUTOPLAY().MUTEDSet(true),
		},
		{
			Expected: "<video autoplay></video>",
			Actual:   VIDEO().AUTOPLAY().MUTEDSet(false),
		},
		{
			Expected: "<video autoplay></video>",
			Actual:   VIDEO().AUTOPLAY().MUTED().MUTEDRemove(),
		},
	})
}

func TestStringAttributes(t *testing.T) {
	run(t, []result{
		{
			Expected: "<button popovertarget=\"my-popover\">Open Popover</button>",
			Actual:   BUTTON().POPOVERTARGET("my-popover").Text("Open Popover"),
		},
	})
}

func TestKVAttributes(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div id=\"elt\" style=\"border-top:1px solid blue;color:red\">An example div</div>",
			Actual:   DIV().ID("elt").STYLE("border-top: 1px solid blue; color: red;").Text("An example div"),
		},
		{
			Expected: "<div style=\"display:none\"></div>",
			Actual:   DIV().STYLEAdd("display", "none"),
		},
		{
			Expected: "<span style=\"color:red;display:block\"></span>",
			Actual: SPAN().
				STYLEAdd("color", "red").
				STYLEMap(map[string]string{"display": "block", "font-size": "12px", "font-weight": "bold"}).
				STYLERemove("font-size", "font-weight"),
		},
		{
			Expected: "<p style=\"display:block;margin:10px;padding:5px\"></p>",
			Actual:   P().STYLEMap(map[string]string{"margin": "10px", "padding": "5px", "display": "block"}),
		},
	})

	assert.NotPanics(t, func() { P().STYLEMap(map[string]string{"foo": ""}) })

	assert.Panics(t, func() { A().STYLEPairs("foo") })
	assert.Panics(t, func() { DIV().STYLEAdd("", "bar") })
	assert.Panics(t, func() { SPAN().STYLE(";;;;;;") })
	assert.Panics(t, func() { DIV().STYLE("font-size; color: red;") })
}

func TestChoiceAttributes(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div id=\"my-popover\" popover=\"auto\">Greetings, one and all!</div>",
			Actual:   DIV().POPOVER(DivPopover_auto).ID("my-popover").Text("Greetings, one and all!"),
		},
		{
			Expected: "<div popover></div>",
			Actual:   DIV().POPOVER(DivPopover_empty),
		},
		{
			Expected: "<a hidden></a>",
			Actual:   A().HIDDEN(AHidden_empty),
		},
		{
			Expected: "<a hidden=\"until-found\"></a>",
			Actual:   A().HIDDEN(AHidden_until_found),
		},
	})
}

func TestSpaceDelimitedAttributes(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div class=\"foo bar baz\"></div>",
			Actual:   DIV().CLASS("foo bar baz hello").CLASSRemove("hello"),
		},
		{
			Expected: "<div class=\"foo bar\"></div>",
			Actual:   DIV().CLASS("foo").CLASS("bar"),
		},
	})
}

func TestCommaDelimitedAttributes(t *testing.T) {
	run(t, []result{
		{
			Expected: "<area alt=\"HTML\" coords=\"260,96,209,249,130,138\" href=\"https://developer.mozilla.org/docs/Web/HTML\" shape=\"poly\">",
			Actual:   AREA().SHAPE(AreaShape_poly).COORDS("260,96,209,249,130,138").HREF("https://developer.mozilla.org/docs/Web/HTML").ALT("HTML"),
		},
		{
			Expected: "<area alt=\"HTML\" coords=\"260,96,209,249,130\" href=\"https://developer.mozilla.org/docs/Web/HTML\" shape=\"poly\">",
			Actual:   AREA().SHAPE(AreaShape_poly).COORDS("260,96,209,249,130,138").COORDSRemove("138").HREF("https://developer.mozilla.org/docs/Web/HTML").ALT("HTML"),
		},
	})
}
