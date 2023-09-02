package pages

import (
	"encoding/json"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type newsPage struct {
	app.Compo
}

func (s *newsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&newsTable{},
	)
}

type newsTable struct {
	app.Compo
	news []NewsItem
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
