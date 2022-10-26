package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

/*
 * PAGES
 */

type homePage struct {
	app.Compo
}

type listPage struct {
	app.Compo
	donors map[string]donor
}

type mapPage struct {
	app.Compo
}

type loginPage struct {
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

func (l *listPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&table{},
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

func (s *loginPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&login{},
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

type table struct {
	app.Compo

	//donors map[string]donor
	donors map[string]User
}

type maps struct {
	app.Compo
}

type welcome struct {
	app.Compo
}

type login struct {
	app.Compo
}

type faq struct {
	app.Compo
}

var navbarCol = "#ed333b"

// top navbar
func (h *header) Render() app.UI {
	return app.Nav().ID("nav").Class("top fixed-top").Style("background-color", navbarCol).Body(
		app.A().Href("/").Text("home").Body(
			app.I().Body(
				app.Text("home")),
			app.Span().Body(
				app.Text("home")),
		),
		app.A().Href("/list").Text("users").Body(
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
		app.A().Href("/login").Text("přihlášení").Body(
			app.I().Body(
				app.Text("login")),
			app.Span().Body(
				app.Text("přihlášení")),
		),
	)
}

type donor struct {
	ID     int
	Name   string
	Active bool
}

func (t *table) OnNav(ctx app.Context) {
	log.Println("starting db read")

	ctx.Async(func() {
		//var donors = make(map[string]donor)
		//var donor donor

		url := "http://swapi.savla.su/users/"

		// push requests use PUT method
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Print(err)
		}

		req.Header.Set("X-Auth-Token", "676c1618a631cffdsf5554xy545n4oo55q33ppvxcx555sa623a5aeea14e42ecfac7e77da8cfbcf4b69d6a3999828e9b0181ade")
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{}

		res, err := client.Do(req)
		defer res.Body.Close()
		if err != nil {
			log.Print(err)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		//var users map[string]interface{}
		var users Users
		if err := json.Unmarshal(data, &users); err != nil {
			log.Print(err)
		}

		//log.Println(len(users))
		log.Println(len(users.Users))

		/*
			donor.ID = 1
			donor.Name = "test"
			donor.Active = true

			donors[string(donor.ID)] = donor
		*/

		// Storing HTTP response in component field:
		ctx.Dispatch(func(ctx app.Context) {
			//log.Println(len(users.users))
			t.donors = users.Users
			log.Println(len(t.donors))

			log.Println("dispatch ends")
		})

		log.Println("db has been read!")
	})
}

func (t *table) Render() app.UI {

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
				app.Range(t.donors).Map(func(k string) app.UI {

					return app.Tr().Body(
						app.Td().Text(t.donors[k].Name),
						app.Td().Text(t.donors[k].FullName),
						app.Td().Text(t.donors[k].Active),
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
		app.Title().Text("úvod"),

		app.Div().Class("space"),
		app.H6().Text("Kapka pro ušáčka"),
		app.Div().Class("space"),

		app.Div().Class("fill medium-height middle-align center-align").Body(
			app.Div().Class("center-align").Body(
				app.I().Class("extra").Text("person"),
				app.H5().Text("You are not following anyone"),
				app.Div().Class("space"),
				app.Nav().Class("no-space").Body(
					app.Div().Class("max field border left-round").Body(
						app.Input().Type("file").Capture("camera").Accept("image/*"),
					),
					app.Button().Class("large right-round").Text("search"),
				),
			),
		),
	)
}

func (l *login) Render() app.UI {
	return app.Div().Body()
}

type QA struct {
	Q string
	A string
}

func (f *faq) Render() app.UI {
	data := map[string]QA{
		"aaa": QA{
			Q: "používá se ketamin k uspávání?",
			A: "doufáme, že ano",
		},
		"aab": QA{
			Q: "kontakt??????",
			A: "ne",
		},
	}

	return app.Main().Class("responsive").Body(

		app.Div().Class("space"),
		app.H6().Text("Často kladené dotazy"),
		app.Div().Class("space"),

		app.Range(data).Map(func(k string) app.UI {
			// simple expansion
			/*
				return app.Details().Body(
					app.Summary().Text(data[k].Q),
					app.P().Text(data[k].A),
				)
			*/

			// custom expansion
			return app.Article().Body(
				app.Details().Body(
					app.Summary().Class("none").Body(
						app.Div().Class("row").Body(
							app.H6().Text(data[k].Q),
							//app.I().Text("arrow_drop_down"),
						),
					),
					app.P().Text(data[k].A),
				),
			)
		}),
		app.Div().Class("space"),
	)
}
