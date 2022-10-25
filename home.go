package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type homePage struct {
	app.Compo
}

type listPage struct {
	app.Compo
}

type mapPage struct {
	app.Compo
}

type statePage struct {
	app.Compo
}

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

type state struct {
	app.Compo
}

type update struct {
	app.Compo
	updateAvailable bool
}

func (h *homePage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&table{},
		&footer{},
	)
}

func (m *mapPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&maps{},
		&footer{},
	)
}

func (l *listPage) Render() app.UI {
	return app.Div()
}

func (s *statePage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&state{},
		&footer{},
	)
}

// top navbar
func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("top fixed-top").Body(
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
	return app.Nav().ID("nav").Class("bottom fixed-bottom").Body(
		app.A().Href("/settings").Text("settings").Body(
			app.I().Body(
				app.Text("settings")),
			app.Span().Body(
				app.Text("settings")),
		),
		app.A().Href("/about").Text("about").Body(
			app.I().Body(
				app.Text("info")),
			app.Span().Body(
				app.Text("info")),
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

	/* use Async()
	reader, err := fetchRemoteStream("http://swapi.savla.su/users")
	err = json.NewDecoder(reader).Decode(data)
	if err != nil {
		panic(err)
	}
	reader.Close()
	*/

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
	coord := ""

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.P().Text("map of usaceks keks"),
		app.Div().Class("large-space"),
		app.IFrame().Class("responsive responsive-iframe").Attr("height", "67%").Attr("frameborder", "0").Attr("scrolling", "no").Attr("marginheight", "0").Attr("marginwidth", "0").
			Src("https://www.openstreetmap.org/export/embed.html?bbox=8.712158203125002%2C47.724544549099676%2C19.984130859375004%2C51.78823192706476&amp;layer=mapnik").Style("border", "1px solid black"),
		app.Div().Class("space"),
		app.Small().Body(
			app.A().Href("https://www.openstreetmap.org/#map=7/49.799/14.348").Text("View Larger Map"),
		),
	)
}

func (s *state) Render() app.UI {
	return app.Div().Class("fill medium-height middle-align center-align").Body(
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
	)
}
