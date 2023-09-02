package pages

import (
	"encoding/json"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type dishPage struct {
	app.Compo

	notificationPermission app.NotificationPermission
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

type dishTable struct {
	app.Compo
	sockets map[string]Socket

	// socket_id to be (un)muted
	muteName string
}

func (d *dishTable) testNotification(ctx app.Context, e app.Event) {
	ctx.Notifications().New(app.Notification{
		Title: "Test",
		Body:  "A test notification",
		Path:  "/dish",
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
