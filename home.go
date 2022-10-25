package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type homePage struct {
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
		// topbar
		&header{},
		&table{},
		// bottombar
		&footer{},
	)
}

func (s *statePage) Render() app.UI {
	return app.Div().Body(
		&header{},
		&state{},
		&footer{},
	)
}

// top navbar
func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("top").Body(
		app.A().Href("/").Text("home").Body(
			app.I().Body(
				app.Text("home")),
			app.Span().Body(
				app.Text("home")),
		),
		app.A().Href("/state").Text("state").Body(
			app.I().Body(
				app.Text("list")),
			app.Span().Body(
				app.Text("list")),
		),
	)
}

// bottom navbar
func (f *footer) Render() app.UI {
	return app.Nav().ID("nav").Class("bottom").Body(
		app.A().Href("/dish").Text("dish").Body(
			app.I().Body(
				app.Text("satellite_uplink")),
			app.Span().Body(
				app.Text("satellite_uplink")),
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
