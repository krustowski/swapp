package main

import (
	"encoding/json"
	"log"

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

	notificationPermission app.NotificationPermission
}

type domainsPage struct {
	app.Compo
}

type depotPage struct {
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
		&footer{},
		&welcome{},
	)
}

func (l *usersPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&usersTable{},
	)
}

func (m *mapPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&mapsRender{},
	)
}

func (d *dishPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&dishTable{},
	)
}

func (d *dishPage) OnMount(ctx app.Context) {
	d.notificationPermission = ctx.Notifications().Permission()
}

func (f *domainsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&domainsTable{},
	)
}

func (f *depotPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&depotTable{},
	)
}

func (f *nodesPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&nodesTable{},
	)
}

func (s *newsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&newsTable{},
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

	// socket_id to be (un)muted
	muteName string
}

type domainsTable struct {
	app.Compo
	domains []Domain
}

type depotTable struct {
	app.Compo
	items []DepotItem
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

/*
 * TABLE AND OTHER RENDERS
 */

func (d *domainsTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var domains Domains

		url := "http://swapi.savla.su/infra/domains"
		data := fetchSWISData(url)
		if data == nil {
			// no nil pointer dereference!
			log.Println("swis data fetch error, nil pointer")
			return
		}

		if err := json.Unmarshal(*data, &domains); err != nil {
			log.Print(err)
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			d.domains = domains.Domains
			log.Println("dispatch ends")
		})
	})
}

func (d *domainsTable) Render() app.UI {

	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.Table().Class("border left-align").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Text("domain fqdn, owner, registrar"),
					app.Th().Text("expiration date"),
				),
			),
			app.TBody().Body(
				app.Range(d.domains).Slice(func(i int) app.UI {
					return app.Tr().Body(
						app.Td().Body(
							app.B().Text(d.domains[i].FQDN).Style("color", "green"),
							app.Br(),
							app.Text(d.domains[i].Owner),
							app.Br(),
							app.Text(d.domains[i].Registrar),
						),
						app.Td().Body(
							app.Text(d.domains[i].Expiration),
						),
					)
				}),
			),
		),
	)
}

func (d *depotTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var depots Depots

		url := "http://swapi.savla.su/depots/krusty"
		data := fetchSWISData(url)
		if data == nil {
			// no nil pointer dereference!
			log.Println("swis data fetch error, nil pointer")
			return
		}

		if err := json.Unmarshal(*data, &depots); err != nil {
			log.Print(err)
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			d.items = depots.Depot.Items
			log.Println("dispatch ends")
		})
	})
}

func (d *depotTable) Render() app.UI {
	return app.Main().Class("responsive").Body(
		app.Div().Class("large-space"),
		app.Table().Class("border left-align").Body(
			app.THead().Body(
				app.Tr().Body(
					app.Th().Text("item desc, misc"),
					app.Th().Text("location"),
				),
			),
			app.TBody().Body(
				app.Range(d.items).Slice(func(i int) app.UI {
					return app.Tr().Body(
						app.Td().Body(
							app.B().Style("color", "green").Text(d.items[i].Desc),
							app.Br(),
							app.Text(d.items[i].Misc),
						),
						app.Td().Body(
							app.Text(d.items[i].Location),
						),
					)
				}),
			),
		),
	)
}

func (d *dishTable) testNotification(ctx app.Context, e app.Event) {
	ctx.Notifications().New(app.Notification{
		Title: "Test",
		Body:  "A test notification",
		Path:  "/mypage",
	})
}

func (d *dishTable) onClick(ctx app.Context, e app.Event) {
	element := ctx.JSSrc().Get("value").String()
	log.Println(element)
}

func (d *dishTable) OnNav(ctx app.Context) {
	queryMute := app.Window().URL().Query().Get("mute")

	if queryMute != "" {
		//ctx.Async(func() {
		url := "http://swapi.savla.su/dish/sockets/" + queryMute + "/mute"

		if ok := putSWISData(url); !ok {
			return
		}
		//})
	}

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

		app.Button().Text("Test Notification").
			OnClick(d.testNotification),

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
								app.A().Href("?mute="+d.sockets[k].ID).Body(
									app.Button().OnClick(d.onClick).Value(d.sockets[k].ID).Class("tertiary").Body(
										app.I().Text("warning"),
										app.Span().Text("off"),
									),
								),
							).Else(
								app.A().Href("?mute="+d.sockets[k].ID).Body(
									app.Button().OnClick(d.onClick).Value(d.sockets[k].ID).Class("primary").Body(
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
