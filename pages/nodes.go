package pages

import (
	"encoding/json"
	"log"

	infra "go.savla.dev/swis/v5/pkg/infra"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type NodesPage struct {
	app.Compo
}

func (n *NodesPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&nodesTable{},
	)
}

type nodesTable struct {
	app.Compo
	nodes []infra.Host
}

func (n *nodesTable) OnNav(ctx app.Context) {
	ctx.Async(func() {
		var nodes infra.Hosts

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
			n.nodes = nodes.Hosts
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
							app.A().Href("http://docs.savla.su/nodes/"+node.HostnameShort).
								Style("color", "green").Body(
								app.B().Text(node.HostnameShort),
							),
							app.Br(),
							app.P().Text(node.HostnameFQDN),
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
