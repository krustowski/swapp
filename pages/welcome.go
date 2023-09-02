package pages

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type welcome struct {
	app.Compo
}

func (w *welcome) Render() app.UI {
	return app.Main().Class("responsive").Body(
		app.Div().Class("space"),

		app.H6().Text("savla.dev link tree"),
		app.Div().Class("space"),

		app.A().Href("https://savla.dev").Style("color", "green").Body(
			app.B().Text("savla.dev Homepage"),
		),
		app.Div().Class("space"),

		app.A().Href("https://github.com/savla-dev").Style("color", "green").Body(
			app.B().Text("savla.dev GitHub Homepage"),
		),

		app.Div().Class("space"),

		app.A().Href("http://docs.savla.su/").Style("color", "green").Body(
			app.B().Text("savla.dev Documentation (intra)"),
		),
		app.Div().Class("space"),

		app.Div().Class("large-divider"),

		app.H6().Text("monitoring and metrics"),
		app.Div().Class("space"),

		app.A().Href("http://grafana.savla.su/dashboards").Style("color", "green").Body(
			app.B().Text("Grafana dashboards (intra)"),
		),
		app.Div().Class("space"),

		app.A().Href("http://prometheus.savla.su/alerts").Style("color", "green").Body(
			app.B().Text("Prometheus alerts (intra)"),
		),
		app.Div().Class("space"),
	)
}
