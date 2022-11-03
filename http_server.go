//go:build !wasm
// +build !wasm

package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func initWASM() {
	app.Route("/", &homePage{})
	app.Route("/users", &usersPage{})
	app.Route("/map", &mapPage{})
	app.Route("/dish", &dishPage{})
	app.Route("/depots", &depotPage{})
	app.Route("/nodes", &nodesPage{})
	app.Route("/news", &newsPage{})

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
