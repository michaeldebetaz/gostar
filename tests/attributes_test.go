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
			Expected: "<button popovertarget=\"my-popover\">Open Popover</button>",
			Actual:   BUTTON().POPOVERTARGET("my-popover").Text("Open Popover"),
		},
		{
			Expected: "<div popover=\"auto\" id=\"my-popover\">Greetings, one and all!</div>",
			Actual:   DIV().POPOVER(DivPopover_auto).ID("my-popover").Text("Greetings, one and all!"),
		},
		{
			Expected: "<div popover></div>",
			Actual:   DIV().POPOVER(DivPopover_empty),
		},
		{
			Expected: "<video autoplay muted></video>",
			Actual:   VIDEO().AUTOPLAY().MUTED(),
		},
	})
}
