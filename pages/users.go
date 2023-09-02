package pages

import (
	"encoding/json"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type usersPage struct {
	app.Compo
}

func (l *usersPage) Render() app.UI {
	return app.Div().Body(
		app.Body().Class("dark"),
		&header{},
		&footer{},
		&usersTable{},
	)
}

type usersTable struct {
	app.Compo
	users map[string]User
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
