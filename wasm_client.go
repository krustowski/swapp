//go:build wasm
// +build wasm

package main

import (
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

func initServer() {}
