package main

import (
	"encoding/json"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

/*
 * PAGES
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

type newsPage struct {
	app.Compo
}

type faqPage struct {
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
		&users{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (m *mapPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&maps{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (s *newsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&news{},
		app.Div().Class("large-space"),
		&footer{},
	)
}

func (f *faqPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&faq{},
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

type users struct {
	app.Compo

	//donors map[string]donor
	users map[string]User
}

type maps struct {
	app.Compo
}

type welcome struct {
	app.Compo
}

type faq struct {
	app.Compo
}

type News struct {
	News []NewsItem `json:"news"`
}

type NewsItem struct {
	Title   string `json:"title"`
	Perex   string `json:"perex"`
	Link    string `json:"link"`
	PubDate string `json:"pub_date"`
}

type news struct {
	app.Compo

	news []NewsItem `json:"news"`
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
		app.A().Href("/faq").Text("faq").Body(
			app.I().Body(
				app.Text("info")),
			app.Span().Body(
				app.Text("faq")),
		),
		app.A().Href("/news").Text("news").Body(
			app.I().Body(
				app.Text("newspaper")),
			app.Span().Body(
				app.Text("news")),
		),
	)
}

func (n *news) OnNav(ctx app.Context) {
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

func (n *news) Render() app.UI {

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

func (u *users) OnNav(ctx app.Context) {
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

func (u *users) Render() app.UI {

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

func (m *maps) Render() app.UI {
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
