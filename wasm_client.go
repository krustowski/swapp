//go:build wasm
// +build wasm

package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func initWASM() {
	app.Route("/", &homePage{})
	app.Route("/users", &usersPage{})
	app.Route("/map", &mapPage{})
	app.Route("/dish", &dishPage{})
	app.Route("/depots", &depotPage{})
	app.Route("/domains", &domainsPage{})
	app.Route("/nodes", &nodesPage{})
	app.Route("/news", &newsPage{})

	app.RunWhenOnBrowser()
}

func initServer() {}
