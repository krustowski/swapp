package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

/*
 * PAGES
 */

type homePage struct {
	app.Compo
}

type listPage struct {
	app.Compo
}

type mapPage struct {
	app.Compo
}

type settingsPage struct {
	app.Compo
}

type faqPage struct {
	app.Compo
}

func (h *homePage) Render() app.UI {
	return app.Div().Body(
		//app.Body().Class("dark"),
		&header{},
		&welcome{},
		&footer{},
	)
}

func (l *listPage) Render() app.UI {
	return app.Div().Body(
		//app.Body().Class("dark"),
		&header{},
		&table{},
		&footer{},
	)
}

func (m *mapPage) Render() app.UI {
	return app.Div().Body(
		//app.Body().Class("dark"),
		&header{},
		&maps{},
		&footer{},
	)
}

func (s *settingsPage) Render() app.UI {
	return app.Div().Body(
		//app.Body().Class("dark"),
		&header{},
		&table{},
		&footer{},
	)
}

func (f *faqPage) Render() app.UI {
	return app.Div().Body(
		//app.Body().Class("dark"),
		&header{},
		&faq{},
		&footer{},
	)
}

/*
 * NESTED
 */

type header struct {
	app.Compo
}

type footer struct {
	app.Compo
}

type table struct {
	app.Compo

	data *users
}

type maps struct {
	app.Compo
}

type welcome struct {
	app.Compo
}

type faq struct {
	app.Compo
}

var navbarCol = "#ed333b"

// top navbar
func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("top fixed-top").Style("background-color", navbarCol).Body(
		app.A().Href("/").Text("home").Body(
			app.I().Body(
				app.Text("home")),
			app.Span().Body(
				app.Text("home")),
		),
		app.A().Href("/list").Text("list").Body(
			app.I().Body(
				app.Text("list")),
			app.Span().Body(
				app.Text("list")),
		),
		app.A().Href("/map").Text("map").Body(
			app.I().Body(
				app.Text("map")),
			app.Span().Body(
				app.Text("map")),
		),
	)
}

// bottom navbar
func (f *footer) Render() app.UI {
	return app.Nav().ID("nav").Class("bottom fixed-bottom").Style("background-color", navbarCol).Body(
		app.A().Href("/settings").Text("settings").Body(
			app.I().Body(
				app.Text("settings")),
			app.Span().Body(
				app.Text("settings")),
		),
		app.A().Href("/faq").Text("faq").Body(
			app.I().Body(
				app.Text("info")),
			app.Span().Body(
				app.Text("faq")),
		),
	)
}

func (t *table) Render() app.UI {
	data := &users{
		users: map[string]user{
			"krusty": user{
				Name:     "krusty",
				FullName: "ks",
				Active:   true,
			},
			"usacek": user{
				Name:     "usacek",
				FullName: "pan kapka",
				Active:   true,
			},
		},
	}

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.P().Text("test text"),
		app.Table().Class("border center-align").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Text("lmao"),
					app.Th().Text("wtf"),
					app.Th().Text("kek"),
				),
			),
			app.TBody().Body(
				app.Range(data.users).Map(func(k string) app.UI {
					//s := fmt.Sprintf("%s: %v", k, t.data[k])

					return app.Tr().Body(
						app.Td().Text(data.users[k].Name),
						app.Td().Text(data.users[k].FullName),
						app.Td().Text(data.users[k].Active),
					)
				}),
			),
		),
	)
}

func (m *maps) Render() app.UI {
	//coord := ""

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.P().Text("map of usaceks keks"),
		app.Div().Class("large-space"),
		app.IFrame().Class("responsive responsive-iframe").Height(350).Attr("frameborder", "0").Attr("scrolling", "no").Attr("marginheight", "0").Attr("marginwidth", "0").
			Src("https://www.openstreetmap.org/export/embed.html?bbox=8.712158203125002%2C47.724544549099676%2C19.984130859375004%2C51.78823192706476&amp;layer=mapnik").Style("border", "1px solid black"),
		app.Div().Class("space"),
		app.Small().Body(
			app.A().Href("https://www.openstreetmap.org/#map=7/49.799/14.348").Text("View Larger Map"),
		),
	)
}

func (w *welcome) Render() app.UI {
	return app.Div().Class("responsive").Body(

		app.Div().Class("space"),
		app.H6().Text("Kapka pro ušáčka"),
		app.Div().Class("space"),

		app.Div().Class("fill medium-height middle-align center-align").Body(
			app.Div().Class("center-align").Body(
				app.I().Class("extra").Text("person"),
				app.H5().Text("You are not following anyone"),
				app.Div().Class("space"),
				app.Nav().Class("no-space").Body(
					app.Div().Class("max field border left-round").Body(
						app.Input(),
					),
					app.Button().Class("large right-round").Text("search"),
				),
			),
		),
	)
}

type QA struct {
	Q string
	A string
}

func (f *faq) Render() app.UI {
	data := map[string]QA{
		"aaa": QA{
			Q: "používá se ketamin k uspávání?",
			A: "doufáme, že ano",
		},
		"aab": QA{
			Q: "kontakt??????",
			A: "ne",
		},
	}

	return app.Main().Class("responsive").Body(

		app.Div().Class("space"),
		app.H6().Text("Často kladené dotazy"),
		app.Div().Class("space"),

		app.Range(data).Map(func(k string) app.UI {
			// simple expansion
			/*
				return app.Details().Body(
					app.Summary().Text(data[k].Q),
					app.P().Text(data[k].A),
				)
			*/

			// custom expansion
			return app.Article().Body(
				app.Details().Body(
					app.Summary().Class("none").Body(
						app.Div().Class("row").Body(
							app.H6().Text(data[k].Q),
							app.I().Text("arrow_drop_down"),
						),
					),
					app.P().Text(data[k].A),
				),
			)
		}),
		app.Div().Class("space"),
		app.A().Href("tel:+420728535909").Text("+420 728 535 909"),
	)
}
