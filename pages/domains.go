package pages

import (
	"encoding/json"
	"log"

	infra "go.savla.dev/swis/v5/pkg/infra"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type DomainsPage struct {
	app.Compo
}

func (f *DomainsPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&domainsTable{},
	)
}

type domainsTable struct {
	app.Compo
	domains []infra.Domain
}

func (d *domainsTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		domains := struct {
			Domains []infra.Domain `json:"domains"`
		}{}

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
							app.Text(d.domains[i].RegistrarName),
						),
						app.Td().Body(
							app.Text(d.domains[i].ExpirationDate),
						),
					)
				}),
			),
		),
	)
}
