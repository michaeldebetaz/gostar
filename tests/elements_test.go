package tests

import (
	"testing"

	. "github.com/delaneyj/gostar/elements"
	"github.com/samber/lo"
)

func TestDivElement(t *testing.T) {
	run(t, []result{
		{
			Expected: `<div></div>`,
			Actual:   DIV(),
		},
		{
			Expected: `<div data-foo="bar"></div>`,
			Actual:   DIV().CustomData("foo", "bar"),
		},
		{
			Expected: `<div data-baz="qux" data-bind-foo="bar"></div>`,
			Actual:   DIV().CustomData("bind-foo", "bar").CustomData("baz", "qux"),
		},
	})
}

func TestNavElement(t *testing.T) {
	run(t, []result{
		{
			Expected: `<nav class="navbar"><ol><li><a href="/">Home</a></li><li><a href="/contact">Contact</a></li><li><a href="/about">About</a></li></ol></nav>`,
			Actual: NAV().CLASS("navbar").Children(
				OL(
					LI(A().HREF("/").Text("Home")),
					LI(A().HREF("/contact").Text("Contact")),
					LI(A().HREF("/about").Text("About")),
				),
			),
		},
	})
}

func TestSVGElement(t *testing.T) {
	run(t, []result{
		{
			Expected: `<clipPath id="clip-path"><rect class="cls-1" height="300" id="Rectangle_73" width="300"></rect></clipPath>`,
			Actual: SVG_CLIPPATH().ID("clip-path").Children(
				SVG_RECT().CLASS("cls-1").ID("Rectangle_73").WIDTH(300).HEIGHT(300),
			),
		},
		{
			Expected: `<linearGradient gradientUnits="objectBoundingBox" id="linear-gradient" x1="0.048" x2="0.963" y1="0.5" y2="0.5"><stop offset="0" stop-color="#000000"></stop><stop offset="1" stop-color="#0E67B4"></stop></linearGradient>`,
			Actual: SVG_LINEARGRADIENT(
				SVG_STOP().OFFSET(0).STOP_COLOR("#000000"),
				SVG_STOP().OFFSET(1).STOP_COLOR("#0E67B4"),
			).
				ID("linear-gradient").
				GRADIENT_UNITS("objectBoundingBox").
				X_1(0.048).Y_1(0.5).
				X_2(0.963).Y_2(0.5),
		},
	})
}

func TestHTMLElement(t *testing.T) {
	run(t, []result{
		{
			Expected: `<html><body><div class="header">Page Header</div><div autocapitalize="off" class="bg-red-200 block" style="font-size:12px">bar</div></body></html>`,
			Actual: HTML(
				BODY(
					DIV().CLASS("header").Text("Page Header"),
					DIV().
						STYLE("color", "rad").
						STYLE("font-size", "12px").
						STYLERemove("color").
						CLASS("block", "bg-red-200", "hidden").
						CLASSRemove("hidden").
						AUTOCAPITALIZE(DivAutocapitalize_off).
						Text("bar"),
				),
			),
		},
	})
}

func TestGrouper(t *testing.T) {
	type User struct {
		FirstName      string
		Email          string
		FavoriteColors []string
		RawContent     string
		EscapedContent string
	}

	type Navigation struct {
		Item string
		Link string
	}

	user := &User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
		RawContent:     "<div><p>Raw Content to be displayed</p></div>",
		EscapedContent: "<div><div><div>Escaped</div></div></div>",
	}

	nav := []*Navigation{
		{
			Item: "Link 1",
			Link: "http://www.mytest.com/",
		}, {
			Item: "Link 2",
			Link: "http://www.mytest.com/",
		}, {
			Item: "Link 3",
			Link: "http://www.mytest.com/",
		},
	}

	header := func(title string) ElementRenderer {
		return HEADER(
			TITLE().TextF("%s's Home Page", title),
			DIV().CLASS("header").Text("Page Header"),
		)
	}

	navigation := func(nav []*Navigation) ElementRenderer {
		return NAV(
			UL(
				Range(nav, func(n *Navigation) ElementRenderer {
					return LI(
						A().HREF(n.Link).Text(n.Item),
					)
				}),
			).CLASS("navigation"),
		)
	}

	footer := func() ElementRenderer {
		return FOOTER(DIV().CLASS("footer").Text("copyright 2016"))
	}

	index := func(u *User, nav []*Navigation, title string) ElementRenderer {
		return Group(
			Text("<!DOCTYPE html>"),
			HTML(
				BODY(
					header(title),
					navigation(nav),
					SECTION(
						DIV(
							DIV(
								H4().TextF("Hello %s", u.FirstName),
								DIV().CLASS("raw").Text(u.RawContent),
								DIV().CLASS("enc").Escaped(u.EscapedContent),
							).CLASS("welcome"),
							Range(lo.Range(5), func(i int) ElementRenderer {
								count := i + 1
								return Tern(
									count == 1,
									P().TextF("%s has %d message", u.FirstName, count),
									P().TextF("%s has %d messages", u.FirstName, count),
								)
							}),
						).CLASS("content"),
					),
					footer(),
				),
			),
		)
	}
	run(t, []result{
		{
			Expected: `<!DOCTYPE html><html><body><header><title>Bob's Home Page</title><div class="header">Page Header</div></header><nav><ul class="navigation"><li><a href="http://www.mytest.com/">Link 1</a></li><li><a href="http://www.mytest.com/">Link 2</a></li><li><a href="http://www.mytest.com/">Link 3</a></li></ul></nav><section><div class="content"><div class="welcome"><h4>Hello Bob</h4><div class="raw"><div><p>Raw Content to be displayed</p></div></div><div class="enc">&lt;div&gt;&lt;div&gt;&lt;div&gt;Escaped&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;</div></div><p>Bob has 1 message</p><p>Bob has 2 messages</p><p>Bob has 3 messages</p><p>Bob has 4 messages</p><p>Bob has 5 messages</p></div></section><footer><div class="footer">copyright 2016</div></footer></body></html>`,
			Actual:   index(user, nav, user.FirstName),
		},
	})
}
