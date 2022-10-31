package main

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

/*
 * PAGE ROOTS
 */

type homePage struct {
	app.Compo
}

type usersPage struct {
	app.Compo
}

type mapPage struct {
	app.Compo
}

type dishPage struct {
	app.Compo
}

type nodesPage struct {
	app.Compo
}

type newsPage struct {
	app.Compo
}

func (h *homePage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&welcome{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (l *usersPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&usersTable{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (m *mapPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&mapsRender{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (f *dishPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&dishTable{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (f *nodesPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&nodesTable{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (s *newsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&newsTable{},
		app.Div().Class("large-space"),
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

type welcome struct {
	app.Compo
}

type usersTable struct {
	app.Compo
	users map[string]User
}

type mapsRender struct {
	app.Compo
}

type dishTable struct {
	app.Compo
	sockets map[string]Socket
}

type nodesTable struct {
	app.Compo
	nodes []Node
}

type newsTable struct {
	app.Compo
	news []NewsItem
}

//var navbarCol = "#ed333b"
var navbarCol = "#006600"

// top navbar
func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("top fixed-top").Style("background-color", navbarCol).Body(
		app.A().Href("/").Text("home").Body(
			app.I().Body(
				app.Text("home")),
			app.Span().Body(
				app.Text("home")),
		),
		app.A().Href("/users").Text("users").Body(
			app.I().Body(
				app.Text("group")),
			app.Span().Body(
				app.Text("users")),
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
		app.A().Href("/dish").Text("dish").Body(
			app.I().Body(
				app.Text("satellite_alt")),
			app.Span().Body(
				app.Text("dish")),
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

/*
 * TABLE AND OTHER RENDERS
 */

func (d *dishTable) OnMount(u *url.URL) {
	socket := u.Query().Get("mute")
	if socket != "" {
		log.Println(socket)
	}
}

func (d *dishTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var sockets Sockets

		url := "http://swapi.savla.su/dish/sockets"
		data := fetchSWISData(url)
		if data == nil {
			// no nil pointer dereference!
			log.Println("swis data fetch error, nil pointer")
			return
		}

		if err := json.Unmarshal(*data, &sockets); err != nil {
			log.Print(err)
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			d.sockets = sockets.Sockets
			log.Println("dispatch ends")
		})
	})
}

func (d *dishTable) Render() app.UI {
	return app.Main().Class("responsive").Body(
		app.Div().Class("space"),
		app.Table().Class("border left-align").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Text("id,name,tcp,path"),
					app.Th().Text("muted"),
				),
			),
			app.TBody().Body(
				app.Range(d.sockets).Map(func(k string) app.UI {

					return app.Tr().Body(
						app.Td().Body(
							app.B().Text(d.sockets[k].ID).Style("color", "green"),
							app.Br(),
							app.Text(d.sockets[k].Hostname),
							app.Text(":"),
							app.Text(d.sockets[k].PortTCP),
							app.Br(),
							app.Text(d.sockets[k].PathHTTP),
						),
						app.Td().Body(
							app.If(d.sockets[k].Muted,
								app.A().Href("/dish?mute="+d.sockets[k].ID).Body(
									app.Button().Class("tertiary responsive").Body(
										app.I().Text("warning"),
										app.Span().Text("off"),
									),
								),
							).Else(
								app.A().Href("/dish?mute="+d.sockets[k].ID).Body(
									app.Button().Class("primary responsive").Body(
										app.I().Text("check"),
										app.Span().Text("on"),
									),
								),
							),
						),
					)
				}),
			),
		),
	)
}

func (n *nodesTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var nodes Nodes

		url := "http://swapi.savla.su/infra/hosts"
		data := fetchSWISData(url)
		if data == nil {
			// no nil pointer dereference!
			log.Println("swis data fetch error, nil pointer")
			return
		}

		if err := json.Unmarshal(*data, &nodes); err != nil {
			log.Print(err)
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			n.nodes = nodes.Nodes
			log.Println("dispatch ends")
		})
	})
}

func (n *nodesTable) Render() app.UI {

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.Table().Class("border left-align").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Text("node name"),
					app.Th().Text("ip addresses"),
				),
			),
			app.TBody().Body(
				app.Range(n.nodes).Slice(func(i int) app.UI {

					node := n.nodes[i]
					return app.Tr().Body(
						app.Td().Body(
							app.A().Href("http://docs.savla.su/nodes/"+node.NameShort).
								Style("color", "green").Body(
								app.B().Text(node.NameShort),
							),
							app.Br(),
							app.P().Text(node.NameFQDN),
						),
						app.Td().Body(
							app.Range(n.nodes[i].IPAddress).Slice(func(j int) app.UI {
								return app.Div().Body(
									app.Text(node.IPAddress[j]),
									app.Br(),
								)
							}),
						),
					)
				}),
			),
		),
	)
}

func (n *newsTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var news News

		url := "http://swapi.savla.su/news/krusty"
		data := fetchSWISData(url)
		if data == nil {
			// no nil pointer dereference!
			log.Println("swis data fetch error, nil pointer")
			return
		}

		if err := json.Unmarshal(*data, &news); err != nil {
			log.Print(err)
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			n.news = news.News
			log.Println("dispatch ends")
		})
	})
}

func (n *newsTable) Render() app.UI {

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.Table().Class("border left-align").Body(

			app.TBody().Body(
				app.Range(n.news).Slice(func(i int) app.UI {

					return app.Tr().Body(
						app.A().Href(n.news[i].Link).Style("color", "green").Body(
							app.B().Text(n.news[i].Title),
						),
						app.Br(),
						app.Small().Text(n.news[i].PubDate),
						app.Br(),
						app.Small().Text(n.news[i].Server),
						app.Div().Class("space"),
						app.Text(n.news[i].Perex),
						app.Div().Class("space"),
						app.Hr(),
						app.Div().Class("space"),
					)
				}),
			),
		),
	)
}

func (u *usersTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var users Users

		url := "http://swapi.savla.su/users/"
		data := fetchSWISData(url)
		if data == nil {
			// no nil pointer dereference!
			log.Println("swis data fetch error, nil pointer")
			return
		}

		if err := json.Unmarshal(*data, &users); err != nil {
			log.Print(err)
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			u.users = users.Users
			log.Println("dispatch ends")
		})
	})
}

func (u *usersTable) Render() app.UI {

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.P().Text("test text"),
		app.Table().Class("border center-align").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Text("id"),
					app.Th().Text("name"),
					app.Th().Text("active"),
				),
			),
			app.TBody().Body(
				//app.Range(data.users).Map(func(k string) app.UI {
				app.Range(u.users).Map(func(k string) app.UI {

					return app.Tr().Body(
						app.Td().Text(u.users[k].Name),
						app.Td().Text(u.users[k].FullName),
						app.Td().Text(u.users[k].Active),
					)
				}),
			),
		),
	)
}

func (m *mapsRender) Render() app.UI {
	coord := "8.712158203125002%2C47.724544549099676%2C19.984130859375004%2C51.78823192706476&amp;"

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.P().Text("map of usaceks keks"),
		app.Div().Class("large-space"),
		app.IFrame().Class("responsive responsive-iframe").Height(350).Attr("frameborder", "0").Attr("scrolling", "no").Attr("marginheight", "0").Attr("marginwidth", "0").
			Src("https://www.openstreetmap.org/export/embed.html?bbox="+coord+"layer=mapnik").Style("border", "1px solid black"),
		app.Div().Class("space"),
		app.Small().Body(
		//app.A().Href("https://www.openstreetmap.org/#map=7/49.799/14.348").Text("View Larger Map"),
		),
	)
}

func (w *welcome) Render() app.UI {
	return app.Main().Class("responsive").Body(
		app.Div().Class("space"),
		app.H6().Text("swAPP home"),
		app.Div().Class("space"),
	)
}
