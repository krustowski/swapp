//go:build !wasm
// +build !wasm

package main

import (
	"log"
	"net/http"

	"swapp/pages"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func initWASM() {
	app.Route("/", &pages.WelcomePage{})
	//app.Route("/login", &pages.LoginPage{})
	app.Route("/users", &pages.UsersPage{})
	app.Route("/map", &pages.MapPage{})
	app.Route("/dish", &pages.DishPage{})
	app.Route("/depots", &pages.DepotPage{})
	app.Route("/domains", &pages.DomainsPage{})
	app.Route("/nodes", &pages.NodesPage{})
	app.Route("/news", &pages.NewsPage{})

	app.RunWhenOnBrowser()
}

func initServer() {
	http.Handle("/", &app.Handler{
		Name:        "swAPP",
		Description: "sakalWeb progressive web app",
		Icon: app.Icon{
			Default:    "/web/logo_284.png",
			AppleTouch: "/web/apple-touch-icon.png",
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/beercss@3.3.3/dist/cdn/beer.min.css",
		},
		Scripts: []string{
			"https://cdn.jsdelivr.net/npm/beercss@3.3.3/dist/cdn/beer.min.js",
			"https://cdn.jsdelivr.net/npm/material-dynamic-colors@1.0.1/dist/cdn/material-dynamic-colors.min.js",
		},
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
