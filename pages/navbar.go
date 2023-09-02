package pages

import "github.com/maxence-charriere/go-app/v9/pkg/app"

//var navbarCol = "#ed333b"
var navbarCol = "#006600"

type header struct {
	app.Compo
}

type footer struct {
	app.Compo
}

// top navbar
func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("top fixed-top").Style("background-color", navbarCol).Body(
		app.A().Href("/").Text("home").Body(
			app.I().Body(
				app.Text("home")),
			app.Span().Body(
				app.Text("home")),
		),
		app.H4().Text("swAPP"),
		app.A().Href("/users").Text("users").Body(
			app.I().Body(
				app.Text("group")),
			app.Span().Body(
				app.Text("users")),
		),
	)
}

// bottom navbar
func (f *footer) Render() app.UI {
	return app.Nav().ID("nav").Class("bottom fixed-bottom").Style("background-color", navbarCol).Body(
		app.A().Href("/dish").Text("dish").Body(
			app.I().Body(
				app.Text("satellite_alt")),
			app.Span().Body(
				app.Text("dish")),
		),
		app.A().Href("/domains").Text("domains").Body(
			app.I().Body(
				app.Text("checklist")),
			app.Span().Body(
				app.Text("domains")),
		),
		app.A().Href("/depots").Text("depots").Body(
			app.I().Body(
				app.Text("inventory")),
			app.Span().Body(
				app.Text("depots")),
		),
		app.A().Href("/nodes").Text("nodes").Body(
			app.I().Body(
				app.Text("dns")),
			app.Span().Body(
				app.Text("nodes")),
		),
		app.A().Href("/news").Text("news").Body(
			app.I().Body(
				app.Text("newspaper")),
			app.Span().Body(
				app.Text("news")),
		),
	)
}
