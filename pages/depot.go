package pages

import (
	"encoding/json"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type depotPage struct {
	app.Compo
}

func (d *depotPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&depotTable{},
	)
}

type depotTable struct {
	app.Compo
	items []DepotItem
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
