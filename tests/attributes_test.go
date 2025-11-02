package tests

import (
	"testing"

	. "github.com/delaneyj/gostar/elements"
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
			Expected: "<div style=\"display:none\"></div>",
			Actual:   DIV().STYLE("display", "none"),
		},
		{
			Expected: "<span style=\"color:red;display:block\"></span>",
			Actual: SPAN().STYLE("color", "red").STYLEMap(map[string]string{
				"display":     "block",
				"font-size":   "12px",
				"font-weight": "bold",
			}).STYLERemove("font-size", "font-weight"),
		},
	})
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
