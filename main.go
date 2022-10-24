package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type home struct {
	app.Compo
}

type header struct {
	app.Compo
}

type table struct {
	app.Compo

	data map[string]obj
}

type obj struct {
	Name  string
	Type  string
	Other string
}

func (h *home) Render() app.UI {
	return app.Div().Body(
		app.Ul().Attr("flex-shrink", 0).Attr("display", "flex"),
		app.Body().Class("dark"),
		// topbar
		&header{},
		// test table
		&table{},
	)
}

func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("m l top surface-variant").Body(
		app.A().Href("/").Text("home").Body(
			app.I().Body(
				app.Text("home")),
			app.Span().Body(
				app.Text("home")),
		),
	)
}

func (t *table) Render() app.UI {
	t.data = map[string]obj{
		"aaaaa": obj{
			Name:  "ooo",
			Type:  "ppp",
			Other: "ok",
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
				app.Range(t.data).Map(func(k string) app.UI {
					//s := fmt.Sprintf("%s: %v", k, t.data[k])

					return app.Tr().Body(
						app.Td().Text(t.data[k].Name),
						app.Td().Text(t.data[k].Type),
						app.Td().Text(t.data[k].Other),
					)
				}),
			),
		),
	)
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &home{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "swAPP",
		Description: "sakalWeb progressive web app",
		Icon: app.Icon{
			Default:    "/web/logo_284.png",
			AppleTouch: "/web/apple-touch-icon.png",
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/beercss@2.3.0/dist/cdn/beer.min.css",
		},
		Scripts: []string{
			"https://cdn.jsdelivr.net/npm/beercss@2.3.0/dist/cdn/beer.min.js",
			"https://cdn.jsdelivr.net/npm/material-dynamic-colors@0.0.10/dist/cdn/material-dynamic-colors.min.js",
		},
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
