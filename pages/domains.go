package pages

import (
	"encoding/json"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type domainsPage struct {
	app.Compo
}

func (f *domainsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&domainsTable{},
	)
}

type domainsTable struct {
	app.Compo
	domains []Domain
}

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
